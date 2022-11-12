package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/go-nlp/distance"
	"github.com/mx79/go-nlp/utils"
)

// WerRoute set up api route for the Word Error Rate calculation
func WerRoute(router fiber.Router) {
	router.Post("", queryWer)
}

// queryWer is the handler func for "post" request to http://localhost:3000/api/v1/wer.
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/wer
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text1\": \"(text...)\", \"text2\": \"(text...)\"}"
func queryWer(c *fiber.Ctx) error {
	var (
		body map[string]string
		res  string
	)

	// Checking X-API-KEY of the request to see if user is allowed or not
	err := checkApiKey(c)
	if err != nil {
		return err
	}

	// Unmarshalling request body before processing it
	err = c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Unable to parse the body of the request",
			"error":   err,
		})
	}

	// Pos Tagging document information
	if utils.MapContains(body, "text1") {
		if utils.MapContains(body, "text2") {
			res = fmt.Sprintf("%v", distance.WordErrorRate(body["text1"], body["text2"]))
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "The \"text2\" parameter is missing in the request body",
				"error":   err,
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "The \"text1\" parameter is missing in the request body",
			"error":   err,
		})
	}

	return c.SendString(res)
}
