package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/vault/api"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnectionHelper() (*mongo.Client, error) {
	// MongoDB connection helper
	uri := VaultConnectionHelper()
	// Placeholder for MongoDB connection and insertion logic
	//timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB connection URI
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	//defer client.Disconnect(ctx)
	client.Database("ClaimDB").Collection("claims")
	return client, nil

}

// CreateClaim godoc
// @Summary      Create a new claim
// @Description  Adds a new claim
// @Tags         claims
// @Accept       json
// @Produce      json
// @Param        claim  body      Claim  true  "Claim to create"
// @Success      201    {object}  Claim
// @Failure      400    {object}  map[string]string "Invalid input"
// @Router       /claims/v1.0 [post]
func SaveClaim(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for save claim logic
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}

	//request to db
	collection := mongoClient.Database("ClaimDB").Collection("claims")
	var claim Claim
	json.NewDecoder(request.Body).Decode(&claim)
	_, err = collection.InsertOne(context.TODO(), claim)
	writer.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(claim)

}

// GetAllClaims godoc
// @Summary      Get all claims
// @Description  Returns list of claims
// @Tags         claims
// @Accept       json
// @Produce      json
// @Success      200  {array}   Claim
// @Router       /claims/v1.0 [get]
func GetClaims(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for get claims logic
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("ClaimDB").Collection("claims")
	cursor, err := collection.Find(context.TODO(), struct{}{})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	var claims []Claim
	if err = cursor.All(context.TODO(), &claims); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(claims)

}

// GetClaimById godoc
// @Summary     Get details of requested claim
// @Description Get details of requested claim
// @Tags        claims
// @Accept      json
// @Produce     json
// @Param       claimid path int true "ID of the Claim"
// @Success     200 {object} Claim
// @Failure     400 {object} map[string]string "Invalid ID supplied"
// @Failure     404 {object} map[string]string "Claim not found"
// @Router      /claims/v1.0/{claimid} [get]
func GetClaimByID(writer http.ResponseWriter, request *http.Request) {

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
	var claim Claim
	if err := collection.FindOne(context.TODO(), bson.M{"claimid": claimID}).Decode(&claim); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(writer, `{"error":"Claim not found"}`, http.StatusNotFound)
			return
		}
		http.Error(writer, `{"error":"`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(claim)

}

// UpdateClaim godoc
// @Summary     Update existing claim
// @Description Update existing claim with the input payload
// @Tags        claims
// @Accept      json
// @Produce     json
// @Param       claimid path int   true "ID of the Claim to update"
// @Param       claim   body Claim true "Updated claim fields"
// @Success     200 {object} Claim
// @Failure     400 {object} map[string]string "Invalid ID supplied"
// @Failure     404 {object} map[string]string "Claim not found"
// @Router      /claims/v1.0/{claimid} [put]
func UpdateClaim(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for delete claim logic
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
	var claim Claim
	json.NewDecoder(request.Body).Decode(&claim)
	_, err = collection.UpdateOne(context.TODO(), bson.M{"claimid": claimID}, bson.M{"$set": claim})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(claim)

}

// DeleteClaimById godoc
// @Summary Delete requested claim
// @Description Delete requested claim
// @Tags claims
// @Accept  json
// @Produce  json
// @Param claimid path int true "ID of the Claim"
// @Success 200 {object} Claim
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Claim not found"
// @Router /claims/v1.0/{claimid} [delete]
func DeleteClaim(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for delete claim logic
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
	_, err = collection.DeleteOne(context.TODO(), map[string]interface{}{"claimid": claimID})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func VaultConnectionHelper() string {
	println("Connecting to vault")
	token := "s.OlqQCVB7IHYx1H30AzAfnZpf"
	vaultAddr := "http://vault:8200"

	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: &http.Client{Timeout: 100 * time.Second}})
	if err != nil {
		println("Error connecting to vault:", err)
	}
	client.SetToken(token)
	println("Connected to vault")
	// If your secrets engine at mount "secret/" is KV v2:
	kv := client.KVv2("secret")

	// name is "mysqlsecret" (the entry you see in UI)
	secret, err := kv.Get(context.Background(), "mongodbpath")
	if err != nil {
		log.Fatalf("vault read: %v", err)
	}
	if secret == nil || secret.Data == nil {
		log.Fatal("vault: empty secret or no data")
	}

	// Keys: mysqlusername, mysqlpassword (from your screenshot)
	uri, _ := secret.Data["uri"].(string)

	fmt.Println("uri:", uri)

	return uri

}
