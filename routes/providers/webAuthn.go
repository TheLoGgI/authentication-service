package providers

// import ("net/http"
// // Wherever you handle your WebAuthn requests
//  "github.com/duo-labs/webauthn/protocol"
//  "github.com/duo-labs/webauthn/webauthn"
// )

// //  https://github.com/duo-labs/webauthn

// // Registering an account
// func BeginRegistration(w http.ResponseWriter, r *http.Request) {
// 	user := datastore.GetUser() // Find or create the new user
// 	options, sessionData, err := web.BeginRegistration(&user)
// 	// handle errors if present
// 	// store the sessionData values
// 	JSONResponse(w, options, http.StatusOK) // return the options generated
// 	// options.publicKey contain our registration options
// }

// func FinishRegistration(w http.ResponseWriter, r *http.Request) {
// 	user := datastore.GetUser() // Get the user
// 	// Get the session data stored from the function above
// 	// using gorilla/sessions it could look like this
// 	sessionData := store.Get(r, "registration-session")
// 	parsedResponse, err := protocol.ParseCredentialCreationResponseBody(r.Body)
// 	credential, err := web.CreateCredential(&user, sessionData, parsedResponse)
// 	// Handle validation or input errors
// 	// If creation was successful, store the credential object
// 	JSONResponse(w, "Registration Success", http.StatusOK) // Handle next steps
// }

// // Logging into an account
// func BeginLogin(w http.ResponseWriter, r *http.Request) {
// 	user := datastore.GetUser() // Find the user
// 	options, sessionData, err := webauthn.BeginLogin(&user)
// 	// handle errors if present
// 	// store the sessionData values
// 	JSONResponse(w, options, http.StatusOK) // return the options generated
// 	// options.publicKey contain our registration options
// }

// func FinishLogin(w http.ResponseWriter, r *http.Request) {
// 	user := datastore.GetUser() // Get the user
// 	// Get the session data stored from the function above
// 	// using gorilla/sessions it could look like this
// 	sessionData := store.Get(r, "login-session")
// 	parsedResponse, err := protocol.ParseCredentialRequestResponseBody(r.Body)
// 	credential, err := webauthn.ValidateLogin(&user, sessionData, parsedResponse)
// 	// Handle validation or input errors
// 	// If login was successful, handle next steps
// 	JSONResponse(w, "Login Success", http.StatusOK)
// }

// // Modifying Credential Options

// var webAuthnHandler webauthn.WebAuthn // init this in your init function

// func beginRegistration() {
//     // Updating the AuthenticatorSelection options.
//     // See the struct declarations for values
//     authSelect := protocol.AuthenticatorSelection{
// 		AuthenticatorAttachment: protocol.AuthenticatorAttachment("platform"),
// 		RequireResidentKey: protocol.ResidentKeyUnrequired(),
//         UserVerification: protocol.VerificationRequired
//     }

//     // Updating the ConveyencePreference options.
//     // See the struct declarations for values
//     conveyancePref := protocol.ConveyancePreference(protocol.PreferNoAttestation)

//     user := datastore.GetUser() // Get the user
//     opts, sessionData, err webAuthnHandler.BeginRegistration(&user, webauthn.WithAuthenticatorSelection(authSelect), webauthn.WithConveyancePreference(conveyancePref))

//     // Handle next steps
// }
