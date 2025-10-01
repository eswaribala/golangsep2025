package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cog/claimapiconsumer/store"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {

	db := store.MySQLConnectionHelper()
	store.GetTableInstance(db)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("Environment variables loaded successfully")
	servername := os.Getenv("servername")
	topic := os.Getenv("topic")
	group := os.Getenv("group")
	log.Printf("Servername: %s, Topic: %s, Group: %s", servername, topic, group)
	// kafka Reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{servername},
		Topic:          topic,
		GroupID:        group, // consumer group for auto rebalancing + offset commits
		MinBytes:       1,     // fetch floor
		MaxBytes:       10e6,  // fetch ceiling (~10MB)
		MaxWait:        500 * time.Millisecond,
		CommitInterval: time.Second, // how often to commit if using ReadMessage
	})
	defer r.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("Consuming topic=%s broker=%s group=%s ...", topic, servername, group)

	for {
		// Use FetchMessage + CommitMessages for at-least-once AFTER processing
		m, err := r.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				log.Println("Shutting down consumer...")
				return
			}
			log.Fatalf("fetch error: %v", err)
		}

		// --- Your processing here ---
		log.Printf("partition=%d offset=%d key=%s value=%s", m.Partition, m.Offset, string(m.Key), string(m.Value))
		//convert m.Value key pair value
		//json to map

		var ClaimData map[string]interface{}
		if err := json.Unmarshal(m.Value, &ClaimData); err != nil {
			fmt.Printf("Failed to parse claim info: %v", err)
			return
		}
		fmt.Println(ClaimData)
		//read key value pair
		for k, v := range ClaimData {
			if k == "claim" {
				fmt.Printf("%s : %v\n", k, v)
				//convert v to json string
				jsonString, err := json.Marshal(v)
				if err != nil {
					fmt.Printf("Failed to convert claim info to json string: %v", err)
					return
				}
				fmt.Printf("Claim JSON: %s\n", jsonString)

				store.SaveClaimInfo(string(jsonString))
			}
		}
		if len(m.Headers) > 0 {
			for _, h := range m.Headers {
				log.Printf("  header %s=%s", h.Key, string(h.Value))
			}
		}
		// ----------------------------

		if err := r.CommitMessages(ctx, m); err != nil {
			log.Printf("commit error: %v", err)
		}
	}

}
