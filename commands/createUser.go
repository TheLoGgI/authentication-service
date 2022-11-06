package commands

import (
	"context"
	"log"
	"time"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"go.mongodb.org/mongo-driver/bson"
)

type defaultUser struct {
	id []byte
}

func CreateUser(newUser models.NewUserAccountRequest) models.User {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	// var newUserWithAuthn {}interface = newUser

	newUserBson, _ := bson.Marshal(newUser)

	cursor, err := collection.InsertOne(ctx, newUserBson)

	if err != nil {
		panic(err)
	}

	userUid := cursor.InsertedID
	log.Printf("New User created with %s \n", userUid)
	var newCreatedUser models.User
	userCursor := collection.FindOne(ctx, bson.D{
		{Key: "uid", Value: userUid},
	})
	userCursor.Decode(newCreatedUser)
	log.Println(newCreatedUser)

	return newCreatedUser
}
