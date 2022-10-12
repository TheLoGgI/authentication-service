package models

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	FirstName string             `bson:"firstName,omitempty"`
	LastName  string             `bson:"lastName,omitempty"`
	Username  string             `bson:"username,omitempty"`
	Email     string             `json:"email" bson:"email, omitempty"`
	Password  string             `json:"password" bson:"password, omitempty"`
	Uid       string             `bson:"uid,omitempty"`
	EntryId   primitive.ObjectID `bson:"_id,omitempty"`
}

type NewUserAccountRequest struct {
	Username string
	Email    string
	Password []byte
	Uid      uuid.UUID
}
