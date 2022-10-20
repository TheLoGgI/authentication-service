package providers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TheLoGgI/queries"
	"golang.org/x/crypto/bcrypt"
)

func CookieAuthLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: " + r.Method)

	// Content Request Headers
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// CORS Headers
	// w.Header().Set("Access-Control-Allow-Credentials", "false")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Request-Method, Origin, Accept")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	// w.Header().Set("Access-Control-Max-Age", "86400")

	// Preflight Request Headers (CORS)
	// w.Header().Set("Access-Control-Request-Headers", "Content-Type, Access-Control-Request-Method, Origin")
	// w.Header().Set("Access-Control-Request-Method", http.MethodPost)
	// w.Header().Set("Origin", "http://localhost:3000")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

	// HTTPS is encrypted
	password := r.FormValue("password")
	email := r.FormValue("email")

	fmt.Println("password: ( " + password + " )")
	fmt.Println("email: ( " + email + " )")

	// Find user
	user, userError := queries.GetUserWithEmail(email)

	fmt.Println("userError: ")
	fmt.Println(userError)
	fmt.Println("------------------------------------")
	fmt.Println("user: ")
	fmt.Println(user)

	if userError != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w)
	}

	hashedPasswordsDidNotMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	fmt.Println("hashedPasswordsDidNotMatch: ")
	fmt.Println(hashedPasswordsDidNotMatch)

	if hashedPasswordsDidNotMatch != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password and username din't match\n"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    fmt.Sprintf("usr:%v", user.Username),
		Expires:  time.Now().Add(5 * time.Minute),
		HttpOnly: true,
		// Path:     "/",
		MaxAge: 60 * 60 * 10,
	})

	w.WriteHeader(http.StatusOK)

	fmt.Println("response: ")
	fmt.Println(w.Header().Values("status"))
	// fmt.Println(jsonResp)
	w.Write([]byte("{\"message\":\"Status All Okay\"}"))

	// fmt.Fprintf(w, "User %v logged in", user.Username)

	// w.Write([]byte("Ending"))
}
