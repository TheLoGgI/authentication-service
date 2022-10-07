package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
// var mongoPassword = os.Getenv("MONGOPASS") // OomqdcOZ5HiNGhlW

type User struct {
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
	Username  string
	Email     string
	Uid       primitive.ObjectID
}

var globalClient *mongo.Client

// func InsertOne(ctx *mongo.Collection, bson bson.D) *mongo.InsertOneResult {
// 	// Insert document
// 	// res, err := ctx.InsertOne(context.Background(), bson)
// 	result, err := coll.InsertOne(
// 		context.TODO(),
// 		bson.D{
// 			{"type", "Masala"},
// 			{"rating", 10},
// 			{"vendor", bson.A{"A", "C"}}
// 		}
// 	)

// 	if err != nil {
// 		// return err
// 		fmt.Println(err)
// 	}

// 	return result
// 	// id := res.InsertedID
// 	// fmt.Printf("InsertId: %s", id)
// 	// return id
// }

func PatchUserSignInToken(ctx *mongo.Collection, bson primitive.M) {

}

func GetAuthToken(token string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var collection = globalClient.Database("salvare").Collection("users")

	filter := bson.D{
		{Key: "token", Value: bson.D{{Key: "$in", Value: bson.A{token}}}},
	}

	result := collection.FindOne(ctx, filter)

	fmt.Println("GetAuthToken Result: ")
	fmt.Println(result)
}

func MongoCollection() *mongo.Collection {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Create a new client and connect to the server

	if globalClient == nil {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
		globalClient = client

		if err != nil {
			fmt.Println("Mongo connection failed")
			panic(err)
		}

		// IIFE
		// defer func() {
		// 	if err = client.Disconnect(context.Background()); err != nil {
		// 		panic(err)
		// 	}
		// }()
	}

	return globalClient.Database("salvare").Collection("users")
}

func GetMongoDatabase() *mongo.Database {

	if globalClient == nil {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
		globalClient = client

		if err != nil {
			fmt.Println("Mongo connection failed")
			panic(err)
		}

		// IIFE
		// defer func() {
		// 	if err = client.Disconnect(context.Background()); err != nil {
		// 		panic(err)
		// 	}
		// }()
	}

	return globalClient.Database("salvare")
}
