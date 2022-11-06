package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

// Can map all the header of a request
//res := make(map[string]string)
//c.Request().Header.VisitAll(func(key, value []byte) {
//	res[string(key)] = string(value)
//})

// checkApiKey is verifying the X-API-KEY provided in header in order to update
// user API quto and check if the provided key exists or not
func checkApiKey(c *fiber.Ctx) error {
	// Reading X-API-KEY from header
	apiKey := string(c.Request().Header.Peek("X-API-KEY"))
	// Check if X-API-KEY is in DB
	user := &bson.M{}
	users := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))
	err := users.FindOne(c.Context(), bson.M{"api_key": apiKey}).Decode(user)
	if err != nil {
		return fiber.NewError(401, "Invalid or wrong API Key")
	}

	return nil
}
