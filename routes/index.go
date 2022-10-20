package routes

import (
	"net/http"

	"github.com/TheLoGgI/models"
	"github.com/TheLoGgI/routes/providers"
	"github.com/TheLoGgI/routes/users"
	"github.com/gorilla/mux"
)

// middleware.EnsureValidToken()(
// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
// 	}),
// )

//

func Providers(server models.Server) {

	// Testing
	server.Router.HandleFunc("/cookie/login", providers.CookieAuthLogin).Methods(http.MethodPost)
	server.Router.Use(mux.CORSMethodMiddleware(server.Router))
	// server.Router.HandleFunc("/providers/refreshAuthToken", middleware.AuthenticationMiddleware(providers.RefreshAuthToken))
	// server.Router.HandleFunc("/authorize/login", middleware.AuthenticationMiddleware(providers.Login))
	// .Methods("POST")

	// WebAuthn
	// server.Router.HandleFunc("/webautn/register", middleware.AuthenticationMiddleware(providers.Login))
}

func Users(server models.Server) {

	server.Router.HandleFunc("/create/user", users.CreateUser).Methods("POST") /* Require HTTPS */
}
