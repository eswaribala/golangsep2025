package models

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//self refrential structure

type Claim struct {
	ClaimID     uint   `json:"id" gorm:"primaryKey"`
	ClaimAmount int    `json:"amount"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}

func CreateMongoDBConnection() (*mongo.Client, error) {

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
	defer client.Disconnect(ctx)
	client.Database("ClaimDB").Collection("claims")
	return client, nil
}

// implement the interface methods
func (c *Claim) Save() int64 {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return 0
	}
	collection := client.Database("ClaimDB").Collection("claims")
	result, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		return 0
	}
	id := result.InsertedID.(int64)
	return id

}

func (c *Claim) GetAll() []*Claim {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return nil
	}

	collection := client.Database("ClaimDB").Collection("claims")
	cursor, err := collection.Find(context.TODO(), map[string]interface{}{})
	if err != nil {
		return nil

	}
	defer cursor.Close(context.TODO())
	var claims []*Claim
	for cursor.Next(context.TODO()) {
		var claim Claim
		if err := cursor.Decode(&claim); err != nil {
			return nil
		}
		claims = append(claims, &claim)
	}

	return claims

}

func (c *Claim) GetByID(id uint) *Claim {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return nil

	}
	collection := client.Database("ClaimDB").Collection("claims")
	var claim Claim
	err = collection.FindOne(context.TODO(), map[string]interface{}{"claimid": id}).Decode(&claim)

	if err != nil {
		return nil
	}

	return &claim

}
func (c *Claim) Update(claimID uint, claimAmount int) *Claim {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return nil
	}

	collection := client.Database("ClaimDB").Collection("claims")
	_, err = collection.UpdateOne(context.TODO(), map[string]interface{}{"claimid": claimID}, map[string]interface{}{"$set": map[string]interface{}{"claimamount": claimAmount}})
	if err != nil {
		return nil
	}

	return c

}
func (c *Claim) Delete(id uint) bool {
	client, err := CreateMongoDBConnection()
	if err != nil {
		return false
	}

	collection := client.Database("ClaimDB").Collection("claims")
	_, err = collection.DeleteOne(context.TODO(), map[string]interface{}{"claimid": id})
	if err != nil {
		return false
	} else {
		return true
	}
}
