package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jdkato/prose/v2"
	"github.com/mx79/go-nlp/utils"
)

// NerRoute set up api route for the Named Entity Recognition service
func NerRoute(router fiber.Router) {
	router.Post("", queryNer)
}

// queryNer is the handler func for "post" request to http://localhost:3000/api/v1/ner.
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/ner
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text\": \"(text...)\"}"
func queryNer(c *fiber.Ctx) error {
	var (
		body map[string]string
		res  []byte
	)
	entMap := make(map[string]string)
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

	// Extracting information or returning empty dict if no one
	if utils.MapContains(body, "text") {
		doc, _ := prose.NewDocument(body["text"])
		for _, ent := range doc.Entities() {
			entMap[ent.Label] = ent.Text
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "The \"text\" parameter is missing in the request body",
			"error":   err,
		})
	}
	res, _ = json.Marshal(entMap)

	return c.Send(res)
}
