package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jonreiter/govader"
	"github.com/mx79/go-nlp/utils"
)

// sentimentAnalyzer is the object responsible for the Sentiment Analysis process.
//
// It holds any necessary methods including PolarityScores
var sentimentAnalyzer = govader.NewSentimentIntensityAnalyzer()

// SentimentRoute set up api route for the Sentiment Analysis service
func SentimentRoute(router fiber.Router) {
	router.Post("", querySentiment)
}

// querySentiment is the handler func for "post" request to http://localhost:3000/api/v1/sentiment.
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/sentiment
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text\": \"(text...)\"}"
func querySentiment(c *fiber.Ctx) error {
	var (
		body map[string]string
		res  []byte
	)
	resMap := make(map[string]float64)

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

	// Sentiment analysis scores of the text
	if utils.MapContains(body, "text") {
		sentiment := sentimentAnalyzer.PolarityScores(body["text"])
		resMap["positive"] = sentiment.Positive
		resMap["neutral"] = sentiment.Neutral
		resMap["negative"] = sentiment.Negative
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "The \"text\" parameter is missing in the request body",
			"error":   err,
		})
	}
	res, _ = json.Marshal(resMap)

	return c.Send(res)
}
