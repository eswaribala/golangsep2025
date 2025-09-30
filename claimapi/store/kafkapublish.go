package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	kafkaTopic   = "claim-events"
	kafkaBroker  = "kafka:9092" // or host.docker.internal:9092 if calling from another container
	publishEvent = "claim_fetched"
)

// Reusable writer (initialize once in init() or main())
var claimWriter *kafka.Writer

func EnsureTopic() error {
	println("Ensuring kafka topic")
	// 1) Dial any broker
	b, err := kafka.Dial("tcp", kafkaBroker)
	if err != nil {
		return fmt.Errorf("dial broker %s: %w", kafkaBroker, err)
	}
	defer b.Close()

	// 2) Ask who the controller is
	ctrlInfo, err := b.Controller()
	if err != nil {
		return fmt.Errorf("get controller: %w", err)
	}
	ctrlAddr := net.JoinHostPort(ctrlInfo.Host, strconv.Itoa(ctrlInfo.Port))

	// 3) Dial controller and set a deadline (replacement for context)
	ctrl, err := kafka.Dial("tcp", ctrlAddr)
	if err != nil {
		return fmt.Errorf("dial controller %s: %w", ctrlAddr, err)
	}
	defer ctrl.Close()
	_ = ctrl.SetDeadline(time.Now().Add(10 * time.Second))

	// 4) Create topic (idempotent: ignore "already exists")
	err = ctrl.CreateTopics(kafka.TopicConfig{
		Topic:             kafkaTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "already exist") {
			log.Printf("[kafka] topic %q already exists", kafkaTopic)
			return nil
		}
		return fmt.Errorf("create topic %q: %w", kafkaTopic, err)
	}
	log.Printf("[kafka] created topic %q (partitions=%d, rf=%d)", kafkaTopic, 1, 1)
	return nil
}
func InitKafkaWriters() {
	claimWriter = &kafka.Writer{
		Addr:                   kafka.TCP(kafkaBroker),
		Topic:                  kafkaTopic,
		Balancer:               &kafka.Hash{},
		RequiredAcks:           kafka.RequireAll,
		BatchTimeout:           100 * time.Millisecond,
		AllowAutoTopicCreation: false, // keep explicit creation for stability
		Logger:                 log.New(os.Stdout, "[kafka] ", log.LstdFlags),
		ErrorLogger:            log.New(os.Stderr, "[kafka-err] ", log.LstdFlags),
	}
}

// Call this on app shutdown
func CloseKafkaWriters() {
	if claimWriter != nil {
		_ = claimWriter.Close()
	}
}

// GPublishClaimInfoByID godoc
// @Summary     Get details of requested claim
// @Description Get details of requested claim
// @Tags        claims
// @Accept      json
// @Produce     json
// @Param       claimid path int true "ID of the Claim"
// @Success     200 {object} Claim
// @Failure     400 {object} map[string]string "Invalid ID supplied"
// @Failure     404 {object} map[string]string "Claim not found"
// @Router      /claims/v1.0/kafka/{claimid} [get]
func PublishClaimInfoByID(writer http.ResponseWriter, request *http.Request) {
	var claim Claim
	idStr := request.PathValue("claimid")
	claimID, err := strconv.Atoi(idStr)
	if err != nil || claimID <= 0 {
		http.Error(writer, `{"error":"Invalid ID supplied"}`, http.StatusBadRequest)
		return
	}
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("ClaimDB").Collection("claims")
	//get claim by id

	if err := collection.FindOne(context.TODO(), bson.M{"claimid": claimID}).Decode(&claim); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(writer, `{"error":"Claim not found"}`, http.StatusNotFound)
			return
		}
		http.Error(writer, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	// 3) Publish to Kafka (best-effort, with timeout)
	go func(c Claim) {
		// encode payload
		body, err := json.Marshal(struct {
			Event string `json:"event"`
			Claim Claim  `json:"claim"`
			TS    string `json:"ts"`
		}{
			Event: publishEvent,
			Claim: c,
			TS:    time.Now().Format(time.RFC3339Nano),
		})
		if err != nil {
			log.Printf("[kafka] marshal error: %v", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := claimWriter.WriteMessages(ctx, kafka.Message{
			Key:   []byte("claim-" + strconv.Itoa(int(c.ClaimID))), // stable partitioning by claim id
			Value: body,
			Headers: []kafka.Header{
				{Key: "source", Value: []byte("sites-api")},
				{Key: "op", Value: []byte("read")},
			},
		}); err != nil {
			log.Printf("[kafka] publish error: %v", err)
		}
	}(claim)

	// 4) Respond to client
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(claim)

}
