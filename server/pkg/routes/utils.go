package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/server/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"os"
)

// Can map all the header of a request
//res := make(map[string]string)
//c.Request().Header.VisitAll(func(key, value []byte) {
//	res[string(key)] = string(value)
//})

// hashPassword mixes the entered password and return its hash
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checkPasswordHash checks and compares the entered password and the related hash of this one
//
// If there is a match, returns true, if no match is found, returns false
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// checkApiKey checks the X-API-KEY provided in the header.
//
// This updates the user's API quota and checks whether the provided key exists or not.
func checkApiKey(c *fiber.Ctx) error {
	// Reading X-API-KEY from header
	apiKey := string(c.Request().Header.Peek("X-API-KEY"))

	// Check if X-API-KEY is in DB
	user := &bson.M{}
	users := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))
	err := users.FindOne(c.Context(), bson.M{"api_key": apiKey}).Decode(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid or wrong API Key",
			"error":   err,
		})
	}

	return nil
}
