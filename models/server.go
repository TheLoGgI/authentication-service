package models

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Database *mongo.Database
	Router   *mux.Router
	// email  EmailSender
}
