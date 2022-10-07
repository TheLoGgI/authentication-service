package queries

import (
	"context"
	"time"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(uid string) models.UserAccount {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	findFilter := bson.D{
		{Key: "_id", Value: uid},
	}
	var foundUser models.UserAccount
	cursor := collection.FindOne(ctx, findFilter)
	cursor.Decode(foundUser)
	err := cursor.Err()

	if err != nil {
		panic(err)
	}

	return foundUser
}
