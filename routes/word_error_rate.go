package routes

import "github.com/gofiber/fiber/v2"

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
//	-d "{\"text\": \"(text...)\"}"
func queryWer(c *fiber.Ctx) error {
	var res []byte
	return c.Send(res)
}
