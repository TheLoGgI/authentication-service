package main

import (
	"fmt"
	"log"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"github.com/TheLoGgI/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const Port string = "3000"

func createServer() models.Server {

	app := fiber.New()
	database := database.GetMongoDatabase()
	// router.Use(mux.CORSMethodMiddleware(router))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3001",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	// src := &http.Server{
	// 	Handler: router,
	// 	Addr:    "127.0.0.1:" + Port,
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	server := models.Server{
		Database: database,
		App:      app,
		// Http:     src,
	}

	return server
}

func main() {

	// Create server
	server := createServer()

	// Routes
	server.App.Static("/", "./static")
	routes.Providers(server)
	// routes.Users(server)

	// Listen for port
	fmt.Printf("Starting server at port " + Port + "\n")

	log.Fatal(server.App.Listen("127.0.0.1:" + Port))
}
