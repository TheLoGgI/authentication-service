package routes

import (
	"net/http"

	"github.com/TheLoGgI/middleware"
	"github.com/TheLoGgI/models"
	"github.com/TheLoGgI/routes/providers"
	"github.com/TheLoGgI/routes/users"
)

func Root(server models.Server) {

	// Hosting static files
	path := http.Dir("../static")
	fileServer := http.FileServer(path)
	server.Router.Handle("/", fileServer)

}

// middleware.EnsureValidToken()(
// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
// 	}),
// )

//

func Providers(server models.Server) {

	server.Router.HandleFunc("/providers/refreshAuthToken", middleware.AuthenticationMiddleware(providers.RefreshAuthToken))
	server.Router.HandleFunc("/authorize/login", middleware.AuthenticationMiddleware(providers.Login))
	// .Methods("POST")
}

func Users(server models.Server) {

	server.Router.HandleFunc("/create/user", users.CreateUser).Methods("POST") /* Require HTTPS */
}
