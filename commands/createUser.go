package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(newUser models.NewUserAccountRequest) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	newUserBson, _ := bson.Marshal(newUser)

	cursor, err := collection.InsertOne(ctx, newUserBson)

	if err != nil {
		panic(err)
	}
	userUid := cursor.InsertedID
	fmt.Printf("new User created with %s", userUid)

}
