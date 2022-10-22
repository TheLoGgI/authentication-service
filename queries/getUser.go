package queries

import (
	"context"
	"fmt"
	"time"

	b64 "encoding/base64"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(uid string) (models.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	var foundUser models.User
	cursor := collection.FindOne(ctx, bson.D{
		{Key: "_id", Value: uid},
	})
	cursor.Decode(&foundUser)
	err := cursor.Err()

	return foundUser, err
}

func GetUserWithEmail(email string) (models.User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	var foundUser models.User
	cursor := collection.FindOne(ctx, bson.D{
		{Key: "email", Value: email},
	})
	cursor.Decode(&foundUser)
	err := cursor.Err()

	// if err != nil {
	// 	fmt.Println("testing for failed user fetch")
	// 	return models.User{}
	// }

	return foundUser, err
}

func CreateSessionCookieForUser(user models.User) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	// GenerateSessionId
	sessionId := fmt.Sprintf("%s:%d:%s", user.Uid.String(), time.Now().Add(time.Minute*5).Unix(), user.Username)
	fmt.Println("sessionID: " + sessionId)
	encryptedSession := b64.StdEncoding.EncodeToString([]byte(sessionId))

	var updatedDocument bson.M
	filter := bson.D{{Key: "_id", Value: user.EntryId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "sessionId", Value: encryptedSession}}}}
	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)

	return encryptedSession, err
}

func RemoveSessionCookieForUser(user models.User) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	// GenerateSessionId
	sessionId := fmt.Sprintf("%s:%d:%s", user.Uid, time.Now().Add(time.Minute*5), user.Username)
	encryptedSession := b64.StdEncoding.EncodeToString([]byte(sessionId))

	var updatedDocument bson.M
	filter := bson.D{{Key: "_id", Value: user.EntryId}}
	update := bson.D{{Key: "$unset", Value: bson.D{{Key: "sessionId", Value: ""}}}}
	err := collection.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)

	return encryptedSession, err
}
