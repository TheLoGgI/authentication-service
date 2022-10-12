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

const Port string = "5000"

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

	server := models.Server{
		Database: database,
		Router:   router,
	}

	return server
}

func main() {

	// var password string = "sometoehrpassworsd"
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	// otherhasedpassword, _ := bcrypt.GenerateFromPassword([]byte("sometoehrpassworsd"), bcrypt.MinCost)

	// fmt.Println(hashedPassword)

	// var is = bcrypt.CompareHashAndPassword(otherhasedpassword, hashedPassword)
	// fmt.Println(is)

	// Create server
	server := createServer()

	server.Router.Use(logClients)

	// Routes
	routes.Providers(server)
	routes.Users(server)

	// Listen for port
	fmt.Printf("Starting server at port " + Port + "\n")
	if err := http.ListenAndServe("127.0.0.1:"+Port, server.Router); err != nil {
		log.Fatal(err)
	}
}
