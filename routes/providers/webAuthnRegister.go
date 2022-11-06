package providers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TheLoGgI/commands"
	"github.com/TheLoGgI/queries"
	"go.mongodb.org/mongo-driver/bson"

	// "github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/protocol"
	"github.com/duo-labs/webauthn/webauthn"
	"github.com/gofiber/fiber/v2"
	// "github.com/duo-labs/webauthn.io/session"
)

// Wherever you handle your WebAuthn requests

// //  https://github.com/duo-labs/webauthn
var (
	web *webauthn.WebAuthn
	err error
)

// store := session.New()

// // Registering an account
func BeginRegistration(c *fiber.Ctx) error {

	web, err = webauthn.New(&webauthn.Config{
		RPDisplayName: "Lasse Aakjær",          // Display Name for your site
		RPID:          "localhost",             // Generally the FQDN for your site
		RPOrigin:      "http://localhost:3001", // The origin URL for WebAuthn requests
		// RPIcon:        "http://localhost/logo.png", // Optional icon URL for your site
	})
	if err != nil {
		fmt.Println(err)
	}

	// Updating the AuthenticatorSelection options.
	// See the struct declarations for values
	authSelect := protocol.AuthenticatorSelection{
		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
		RequireResidentKey:      protocol.ResidentKeyUnrequired(),
		UserVerification:        protocol.VerificationPreferred,
	}

	// Updating the ConveyencePreference options.
	// See the struct declarations for values
	conveyancePref := protocol.ConveyancePreference(protocol.PreferNoAttestation)

	// user := datastore.GetUser() // Get the user
	user, _ := queries.GetUserWithEmail("lasse@hotmail.com")

	// generate PublicKeyCredentialCreationOptions, session data
	var opts, sessionData, err = web.BeginRegistration(&user, webauthn.WithAuthenticatorSelection(authSelect), webauthn.WithConveyancePreference(conveyancePref))
	// Get userUid from header or cookie -> c.Cookies("userUid")

	// handle errors if present
	if err != nil {
		fmt.Printf("Failed to register - client: %v \n", err)
	}
	fmt.Printf("Began registration for use %v \n", user.Username)

	// store the sessionData values
	commands.UpdateUser(user.Uid, bson.D{
		{Key: "$set", Value: bson.D{{Key: "session", Value: sessionData}}},
	})

	fmt.Println("Updated session for user.")

	c.SendStatus(http.StatusOK)
	return c.JSON(fiber.Map{
		"options":     &opts,
		"sessionData": &sessionData,
		"userUid":     user.Uid,
		"message":     "Registering user with WebAuthn bio Authentication",
	})
}

func JSONResponse(w http.ResponseWriter, d interface{}, c int) {
	dj, err := json.Marshal(d)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

// No working
func FinishRegistration(w http.ResponseWriter, r *http.Request) {

	// TODO: COPY RESPONSE BODY
	var userUid = "562e9fee-66e5-44c6-b625-5150c8c3368d"
	// session := r.FormValue("session")

	fmt.Println("userUid")
	fmt.Println(userUid)

	user, failedFetchUser := queries.GetUser(userUid)
	if failedFetchUser != nil {
		w.Write([]byte("{\"message\": \"Registration was Unsuccessful\" }"))
		return
	}
	fmt.Println(user)

	// Get the session data stored from the function above
	sessionData := user.Session

	fmt.Println("sessionData")
	fmt.Println(sessionData)
	// sessionData := store.Get(r, "registration-session")
	parsedResponse, parsedBodyError := protocol.ParseCredentialCreationResponseBody(r.Body)

	if parsedBodyError != nil {
		log.Println("failed finishing registration")
		log.Println(parsedBodyError)
		JSONResponse(w, parsedBodyError.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("parsedResponse")
	fmt.Println(parsedResponse)

	credential, credentialError := web.CreateCredential(&user, sessionData, parsedResponse)

	if credentialError != nil {
		log.Println("failed finishing registration")
		log.Println(credentialError)
		JSONResponse(w, credentialError.Error(), http.StatusBadRequest)
		return
	}

	user.AddCredential(credential)

	fmt.Println("credential")
	fmt.Println(credential)

	commands.UpdateUser(user.Uid, bson.D{
		{Key: "$set", Value: bson.D{{Key: "credential", Value: credential}}},
		{Key: "$push", Value: bson.D{{Key: "credentials", Value: credential}}},
	})
	// If creation was successful, store the credential object
	// commands.UpdateUser(user.Uid, bson.D{
	// 	{Key: "$set", Value: bson.D{{Key: "credential", Value: credential}}},
	// 	{Key: "$push", Value: bson.D{{Key: "credentials", Value: credential}}},
	// })

	JSONResponse(w, "Registration was Successful", http.StatusOK)
}

type CredentialCreationResponseData struct {
}

// type CredentialCreationData struct {
// 	Id       string `json:"id" xml:"id" form:"id"`
// 	Type     string `json:"type" xml:"type" form:"type"`
// 	RawID    string `json:"rawId" xml:"rawId" form:"rawId"`
// 	Response protocol.ParsedAttestationResponse
// }

type CredentialCreationData struct {
	UserUid  string `json:"userUid" xml:"userUid" form:"userUid"`
	ID       string `json:"id" xml:"id" form:"id"`
	RawID    string `json:"rawId" xml:"rawId" form:"rawId"`
	Type     string `json:"type" xml:"type" form:"type"`
	Response string `json:"response" xml:"response" form:"response"`
}

// func FinishRegistration(c *fiber.Ctx) error {
// 	// https://github.com/duo-labs/webauthn/blob/master/webauthn/registration.go
// 		fmt.Println("-----------------------FinishRegistration------------------------------")

// 		credentials := new(CredentialCreationData)
// 		c.BodyParser(&credentials)
// 		fmt.Println("body")
// 		fmt.Println(credentials)
// 		return c.SendStatus(http.StatusOK)
// 		// userUid := c.FormValue("userUid")
// 		var userUid = "562e9fee-66e5-44c6-b625-5150c8c3368d"
// 		// session := r.FormValue("session")

// 		fmt.Println("userUid")
// 		fmt.Println(userUid)
// 		// fmt.Println(session) /* [object Object] */
// 		// userUid := "562e9fee-66e5-44c6-b625-5150c8c3368d"

// 		user, failedFetchUser := queries.GetUser(userUid)
// 		if failedFetchUser != nil {
// 			c.SendStatus(http.StatusBadRequest)
// 			return c.JSON(fiber.Map{
// 				"message": "Registration was Unsuccessful",
// 			})
// 		}
// 		fmt.Println(user)

// 		// Get the session data stored from the function above
// 		sessionData := user.Session

// 		fmt.Println("sessionData")
// 		fmt.Println(sessionData)
// 	// https://www.herbie.dev/blog/webauthn-basic-web-client-server/
// 		parsedResponse, err := protocol.ParseCredentialCreationResponse(c.Response())
// 		// sessionData := store.Get(r, "registration-session")
// 		// parsedResponse, _ := protocol.ParseCredentialCreationResponseBody(io.Reader(c.Body()))
// 		// fmt.Println(&parsedResponse)
// 		// fmt.Println(parsedResponse)
// 		// credential, err := web.CreateCredential(&user, sessionData, parsedResponse)
// 		// web.FinishRegistration(user, sessionData, http.Handler(c.Response()))
// 		// user.AddCredential(*credential)

// 		// // Handle validation or input errors
// 		// if err != nil {
// 		// 	fmt.Printf("Failed to create credential: %v \n", err)
// 		// }
// 		// fmt.Println("credential")
// 		// fmt.Println(credential)

// 		// If creation was successful, store the credential object
// 		// commands.UpdateUser(user.Uid, bson.D{
// 		// 	{Key: "$set", Value: bson.D{{Key: "credential", Value: credential}}},
// 		// })

// 		// c.SendStatus(http.StatusOK)
// 		// return c.JSON(fiber.Map{
// 		// 	"message": "Registration Success",
// 		// })

// 		//
// 		c.SendStatus(http.StatusOK)
// 		return c.JSON(fiber.Map{
// 			"message": "Registration Success",
// 		})
// 	}
