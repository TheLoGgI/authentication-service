package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(user models.User) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	updateUserBson, _ := bson.Marshal(user)

	cursor, err := collection.UpdateByID(ctx, user.Uid, updateUserBson)

	if err != nil {
		panic(err)
	}
	userUid := cursor.UpsertedID
	fmt.Printf("User with Uid: %s was updated", userUid)

}
