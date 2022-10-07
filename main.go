package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"github.com/TheLoGgI/routes"
	"github.com/gorilla/mux"
)

const Port string = "3000"

func createServer() models.Server {

	router := mux.NewRouter()
	database := database.GetMongoDatabase()

	server := models.Server{
		Database: database,
		Router:   router,
	}

	return server
}

func main() {

	// Create server
	server := createServer()

	// Routes
	routes.Providers(server)
	routes.Users(server)

	// Listen for port
	fmt.Printf("Starting server at port " + Port + "\n")
	if err := http.ListenAndServe("127.0.0.1:"+Port, server.Router); err != nil {
		log.Fatal(err)
	}
}
