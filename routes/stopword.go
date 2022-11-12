package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/go-nlp/base"
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/utils"
	"github.com/rylans/getlang"
)

// StopwordRoute set up api route for the Stopword service
func StopwordRoute(route fiber.Router) {
	route.Post("", queryStopword)
}

// queryStopword is the handler func for "post" request to http://localhost:3000/api/v1/stopword.
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/stopword
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text\": \"(text...)\"}"
func queryStopword(c *fiber.Ctx) error {
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
		return err
	}

	// Remove stopword in the text if possible
	if utils.MapContains(body, "text") {
		info := getlang.FromString(body["text"])
		stemmer := clean.NewStopwords(base.Lang(info.LanguageCode()))
		res = stemmer.Stop(body["text"])
	} else {
		return fiber.NewError(400, "The \"text\" parameter is missing in the request body")
	}

	return c.SendString(res)
}
