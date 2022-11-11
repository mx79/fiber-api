package routes

import (
	"encoding/json"
	rake "github.com/afjoseph/RAKE.Go"
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/go-nlp/utils"
)

// RakeRoute set up api route for the Rake Algorithm service
func RakeRoute(route fiber.Router) {
	route.Post("", queryRake)
}

// queryRake is the handler func for "post" request to http://localhost:3000/api/v1/rake.
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/rake
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text\": \"(text...)\"}"
func queryRake(c *fiber.Ctx) error {
	body := make(map[string]string)
	keywordsMap := make(map[string]float64)

	// Checking X-API-KEY of the request to see if user is allowed or not
	err := checkApiKey(c)
	if err != nil {
		return err
	}

	// Unmarshalling request body before processing it
	err = c.BodyParser(&body)
	if err != nil {
		return err
	}

	// Extracting information or returning empty dict if no one
	if utils.MapContains(body, "text") {
		keywords := rake.RunRake(body["text"])
		for _, keyword := range keywords {
			keywordsMap[keyword.Key] = keyword.Value
		}
	}
	res, _ := json.Marshal(keywordsMap)

	return c.Send(res)
}
