package users

import (
	"fmt"
	"net/http"

	"github.com/TheLoGgI/commands"
	"github.com/TheLoGgI/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// Check body for password
	password := r.FormValue("password")
	username := r.FormValue("username")
	email := r.FormValue("email")

	responseBody := r.Body

	headers := r.Header
	fmt.Println(responseBody)
	fmt.Printf("AuthToken from client: %s created with password %s \n", headers.Get("Auth-Token"), password)

	// encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	newUser := models.NewUserAccountRequest{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Uid:      "someRandomUid123",
	}
	commands.CreateUser(newUser)
	fmt.Println(newUser)

	w.Write([]byte(`{"message":"User Created"}`))

}
