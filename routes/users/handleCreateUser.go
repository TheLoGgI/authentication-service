package users

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/TheLoGgI/commands"
	"github.com/TheLoGgI/database"
	"github.com/TheLoGgI/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	// TODO: Check registration Methods Google, Github, Facebook, Login Creds
	// authentication ID

	// Check body for password
	password := r.FormValue("password")
	username := r.FormValue("username")
	email := r.FormValue("email")

	// Check email registration
	var foundEmailUser models.User
	collection.FindOne(ctx, bson.D{
		{Key: "email", Value: strings.TrimSpace(email)},
	}).Decode(&foundEmailUser)

	if (foundEmailUser != models.User{}) {
		errMsg := fmt.Sprintf("User with email already exists: %s", foundEmailUser.Username)
		http.Error(w, errMsg, http.StatusBadRequest)
		w.Write([]byte(`{"error" : "400", "message":"User was not Created"}`))
		return
	}

	// Create hashed password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	newUser := models.NewUserAccountRequest{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Uid:      uuid.New(),
	}

	// Create User in database
	commands.CreateUser(newUser)

	w.Write([]byte(`{"message":"User Created"}`))

}

// encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))

// responseBody := r.Body

// headers := r.Header
// fmt.Println(responseBody)
// fmt.Printf("AuthToken from client: %s created with password %s \n", headers.Get("Auth-Token"), password)
