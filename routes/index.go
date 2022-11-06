package routes

import (
	"net/http"

	"github.com/TheLoGgI/models"
	"github.com/TheLoGgI/routes/providers"
	"github.com/TheLoGgI/routes/users"
	"github.com/gofiber/adaptor/v2"
	// "github.com/gofiber/adaptor/v2"
)

// middleware.EnsureValidToken()(
// 	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message":"Hello from a private endpoint! You need to be authenticated to see this."}`))
// 	}),
// )

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func Providers(server models.Server) {

	server.App.Post("/cookie/login", providers.CookieAuthLogin)
	server.App.Get("/webauthn/beginregister", providers.BeginRegistration)
	server.App.Post("/webauthn/finishregistration", adaptor.HTTPHandler(handler(providers.FinishRegistration)))

	server.App.Post("/webauthn/beginlogin", providers.BeginLogin)
	server.App.Post("/webauthn/finishlogin", adaptor.HTTPHandler(handler(providers.FinishLogin)))
	// server.App.Post("/webauthn/finishregistration", adaptor.HTTPHandler(handler(providers.FinishingRegistration)))

	// server.Router.Use(mux.CORSMethodMiddleware(server.Router))
	// server.Router.HandleFunc("/providers/refreshAuthToken", middleware.AuthenticationMiddleware(providers.RefreshAuthToken))
	// server.Router.HandleFunc("/authorize/login", middleware.AuthenticationMiddleware(providers.Login))
	// .Methods("POST")

	// WebAuthn
	// server.Router.HandleFunc("/webautn/register", middleware.AuthenticationMiddleware(providers.Login))
}

func Users(server models.Server) {

	server.App.Post("/create/user", users.CreateUser) /* Require HTTPS */
}
