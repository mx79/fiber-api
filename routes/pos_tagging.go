package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jdkato/prose/v2"
	"github.com/mx79/go-nlp/utils"
)

// PosTaggingRoute set up api route for the Pos Tagging service
func PosTaggingRoute(route fiber.Router) {
	route.Post("", queryPosTagging)
}

// queryPosTagging is the handler func for "post" request to http://localhost:3000/api/v1/pos-tagging.
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/pos-tagging
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text\": \"(text...)\"}"
func queryPosTagging(c *fiber.Ctx) error {
	var (
		body map[string]string
		res  []byte
	)
	resMap := make(map[string]string)

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

	// Pos Tagging document information
	if utils.MapContains(body, "text") {
		doc, _ := prose.NewDocument(body["text"])
		for _, tok := range doc.Tokens() {
			resMap[tok.Tag] = tok.Text
		}
	} else {
		return fiber.NewError(400, "Missing parameter text in request body")
	}
	res, _ = json.Marshal(resMap)

	return c.Send(res)
}
