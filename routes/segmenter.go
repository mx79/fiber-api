package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jdkato/prose/v2"
	"github.com/mx79/go-nlp/utils"
)

// SegmenterRoute
func SegmenterRoute(route fiber.Router) {
	route.Post("", QuerySegmenter)
}

// QuerySegmenter is the handler func for "post" request to http://localhost:3000/api/v1/segmenter,
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/segmenter
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text\": \"(text...)\"}"
func QuerySegmenter(c *fiber.Ctx) error {
	var (
		body          map[string]string
		extractedData []string
		res           []byte
	)

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

	// Segmenting document information
	if utils.MapContains(body, "text") {
		doc, _ := prose.NewDocument(body["text"])
		sents := doc.Sentences()
		for _, sent := range sents {
			extractedData = append(extractedData, sent.Text)
		}
	} else {
		return fiber.NewError(400, "The \"text\" parameter is missing in the request body")
	}
	res, _ = json.Marshal(extractedData)

	return c.Send(res)
}
