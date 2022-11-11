package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/go-nlp/base"
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/utils"
	"github.com/rylans/getlang"
)

// StemmerRoute set up api route for the Stemming service
func StemmerRoute(route fiber.Router) {
	route.Post("", queryStemmer)
}

// queryStemmer is the handler func for "post" request to http://localhost:3000/api/v1/stemmer.
//
// This endpoint can be tested like this:
//
// curl -X POST http://localhost:3000/api/v1/stemmer
// -H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
// -H "Content-Type: application/json"
// -d "{\"text\": \"(text...)\"}"
func queryStemmer(c *fiber.Ctx) error {
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

	// Extracting information or returning empty dict if no one
	if utils.MapContains(body, "text") {
		info := getlang.FromString(body["text"])
		stemmer := clean.NewStemmer(base.Lang(info.LanguageCode()))
		res = stemmer.Stem(body["text"])
	}

	return c.SendString(res)
}
