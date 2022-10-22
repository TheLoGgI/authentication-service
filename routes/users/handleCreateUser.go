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
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.MongoCollection()

	// TODO: Check registration Methods Google, Github, Facebook, Login Creds
	// authentication ID

	// Check body for password
	password := c.FormValue("password")
	username := c.FormValue("username")
	email := c.FormValue("email")

	// Check email registration
	var foundEmailUser models.User
	collection.FindOne(ctx, bson.D{
		{Key: "email", Value: strings.TrimSpace(email)},
	}).Decode(&foundEmailUser)

	if (foundEmailUser != models.User{}) {
		errMsg := fmt.Sprintf("User with email already exists: %s", foundEmailUser.Username)

		c.SendStatus(http.StatusBadRequest)
		return c.SendString(errMsg)
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

	return c.JSON(fiber.Map{
		"message": "User Created",
	})

}

// encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))

// responseBody := r.Body

// headers := r.Header
// fmt.Println(responseBody)
// fmt.Printf("AuthToken from client: %s created with password %s \n", headers.Get("Auth-Token"), password)
