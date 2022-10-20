package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"github.com/TheLoGgI/routes"
	"github.com/gorilla/mux"
)

const Port string = "3000"

// for global use (using a http.Handler!) - https://gist.github.com/AxelRHD/2344cc1105afc06723b363f21486dec8
func logClients(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)

		// compare the return-value to the authMW
		next.ServeHTTP(w, r)
	})
}

func createServer() models.Server {

	router := mux.NewRouter()
	database := database.GetMongoDatabase()
	// router.Use(mux.CORSMethodMiddleware(router))

	src := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	server := models.Server{
		Database: database,
		Router:   router,
		Http:     src,
	}

	return server
}

func Root(server models.Server) {

	// Hosting static files
	// This will serve files under http://localhost:Port/<filename>
	server.Router.Handle("/", http.FileServer(http.Dir("static")))

}

func main() {

	// Create server
	server := createServer()

	// Routes
	Root(server)
	routes.Providers(server)
	routes.Users(server)

	// Listen for port
	fmt.Printf("Starting server at port " + Port + "\n")

	log.Fatal(server.Http.ListenAndServe())
}
