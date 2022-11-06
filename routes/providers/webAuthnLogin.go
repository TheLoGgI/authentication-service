package providers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TheLoGgI/commands"
	"github.com/TheLoGgI/queries"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// var store = session.New()

// // Logging into an account
func BeginLogin(c *fiber.Ctx) error {

	password := c.FormValue("password")
	email := c.FormValue("email")

	fmt.Println(password)
	fmt.Println(email)

	// Find user trying to login
	user, userError := queries.GetUserWithEmail(email)
	fmt.Println(user)

	if userError != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	options, sessionData, loginError := web.BeginLogin(&user)

	// store the sessionData values
	// session.Set(user.Uid.String(), sessionData)
	if loginError != nil {
		fmt.Println("loginError")
		fmt.Println(loginError)
		return c.SendStatus(http.StatusNotFound)
	}

	fmt.Println("sessionData")
	fmt.Println(sessionData)

	commands.UpdateUser(user.Uid, bson.D{
		{Key: "$set", Value: bson.D{{Key: "session", Value: sessionData}}},
	})

	c.SendStatus(http.StatusOK)
	return c.JSON(fiber.Map{
		"options": &options,
		"userUid": user.Uid,
		// "sessionData": &sessionData,
		"status": http.StatusOK,
	})
	// options.publicKey contain our registration options
}

type FinishLoginParams struct {
	Uid string `json:"userUid" form:"userUid"`
}

func FinishLogin(w http.ResponseWriter, r *http.Request) {

	// get username
	// formData := new(CreateUserParams)
	// c.BodyParser(&formData)
	// Declare a new Person struct.
	var p FinishLoginParams

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(p)

	// var userUid = "562e9fee-66e5-44c6-b625-5150c8c3368d"
	// vars := mux.Vars(r)
	// username := vars["username"]

	// get user
	user, getUserError := queries.GetUser(p.Uid)
	// user, err := userDB.GetUser(username)

	// user doesn't exist
	if getUserError != nil {
		log.Println(getUserError)
		JSONResponse(w, getUserError.Error(), http.StatusBadRequest)
		return
	}

	sessionData := user.Session

	// load the session data
	// var sessionData, sessionError = store.Get(userUid)
	// sessionData, err := sessionStore.GetWebauthnSession("authentication", r)
	// if sessionError != nil {
	// 	log.Println(sessionError)
	// 	JSONResponse(w, sessionError.Error(), http.StatusBadRequest)
	// 	return
	// }

	// in an actual implementation we should perform additional
	// checks on the returned 'credential'
	var credentials, verificationError = web.FinishLogin(user, sessionData, r)
	if verificationError != nil {
		log.Println(verificationError)
		JSONResponse(w, verificationError.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(credentials)

	// handle successful login
	JSONResponse(w, "Login Success", http.StatusOK)
}

// func FinishLogin(w http.ResponseWriter, r *http.Request) {
// 	user := datastore.GetUser() // Get the user
// 	// Get the session data stored from the function above
// 	// using gorilla/sessions it could look like this
// 	sessionData := store.Get(r, "login-session")
// 	parsedResponse, err := protocol.ParseCredentialRequestResponseBody(r.Body)
// 	credential, err := web.ValidateLogin(&user, sessionData, parsedResponse)
// 	// Handle validation or input errors
// 	// If login was successful, handle next steps
// 	JSONResponse(w, "Login Success", http.StatusOK)
// }

// func BeginLogin(w http.ResponseWriter, r *http.Request) {
// 	user := datastore.GetUser() // Find the user
// 	options, sessionData, err := web.BeginLogin(&user)
// 	// handle errors if present
// 	// store the sessionData values
// 	JSONResponse(w, options, http.StatusOK) // return the options generated
// 	// options.publicKey contain our registration options
// }
