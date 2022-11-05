package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

// Can map all the header of a request
//res := make(map[string]string)
//c.Request().Header.VisitAll(func(key, value []byte) {
//	res[string(key)] = string(value)
//})

func checkApiKey(c *fiber.Ctx) error {
	// Reading X-API-KEY from header
	res := string(c.Request().Header.Peek("X-API-KEY"))
	// Check if X-API-KEY is in DB
	users := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))
	apiKey, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", res))
	if err != nil {
		return err
	}
	var user bson.M
	err = users.FindOne(c.Context(), bson.M{"_id": apiKey}).Decode(&user)
	if err != nil {
		return fiber.NewError(404, "Invalid or wrong API Key")
	}

	return nil
}
