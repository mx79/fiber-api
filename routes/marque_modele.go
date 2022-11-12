package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/go-nlp/extractor"
	"github.com/mx79/go-nlp/utils"
)

// Loading extractor for brand and model of vehicle
var mmExtractor = extractor.NewLookupExtractor("./resources/marque_modele.json", extractor.IGNORECASE)

// MarqueModeleRoute set up api route for the extractor of brand and model of vehicle
func MarqueModeleRoute(route fiber.Router) {
	route.Post("", queryMarqueModele)
}

// queryMarqueModele is the handler func for "post" request to http://localhost:3000/api/v1/marque-modele.
//
// This endpoint can be tested like this:
//
//	curl -X POST http://localhost:3000/api/v1/marque-modele
//	-H "X-API-KEY: e6b087d4-0c9d-4043-9a72-ffe734811471"
//	-H "Content-Type: application/json"
//	-d "{\"text\": \"(text...)\"}"
func queryMarqueModele(c *fiber.Ctx) error {
	var (
		body          map[string]string
		extractedData map[string]interface{}
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

	// Extracting information or returning empty dict if no one
	if utils.MapContains(body, "text") {
		extractedData = mmExtractor.GetEntity(body["text"])
	} else {
		return fiber.NewError(400, "Missing parameter text in request body")
	}
	res, _ = json.Marshal(extractedData)

	return c.Send(res)
}
