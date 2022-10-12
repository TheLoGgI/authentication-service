package queries

import (
	"context"
	"time"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(uid string) models.User {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	var foundUser models.User
	cursor := collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: uid},
	})
	cursor.Decode(foundUser)
	err := cursor.Err()

	if err != nil {
		panic(err)
	}

	return foundUser
}
