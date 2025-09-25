package store

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnectionHelper() (*mongo.Client, error) {
	// MongoDB connection helper
	_ = godotenv.Load(".env")
	uri := os.Getenv("uri")
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

// crud
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

func GetClaimByID(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for get claim by ID logic
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("ClaimDB").Collection("claims")
	//get claim by id
	var claim Claim
	id := request.URL.Query().Get("id")
	err = collection.FindOne(context.TODO(), map[string]interface{}{"id": id}).Decode(&claim)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(claim)

}

func UpdateClaim(writer http.ResponseWriter, request *http.Request) {
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("ClaimDB").Collection("claims")
	var claim Claim
	json.NewDecoder(request.Body).Decode(&claim)
	_, err = collection.UpdateOne(context.TODO(), map[string]interface{}{"id": claim.ID}, map[string]interface{}{"$set": claim})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(claim)

}

func DeleteClaim(writer http.ResponseWriter, request *http.Request) {
	// Placeholder for delete claim logic
	mongoClient, err := MongoDBConnectionHelper()
	if err != nil {
		writer.Write([]byte("Error connecting to MongoDB"))
		return
	}
	collection := mongoClient.Database("ClaimDB").Collection("claims")
	id := request.URL.Query().Get("id")
	_, err = collection.DeleteOne(context.TODO(), map[string]interface{}{"id": id})
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
