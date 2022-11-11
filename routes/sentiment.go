package routes

import "github.com/gofiber/fiber/v2"

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
	var res []byte
	return c.Send(res)
}
