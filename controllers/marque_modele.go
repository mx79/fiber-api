package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/go-nlp/extractor"
	"github.com/mx79/go-nlp/utils"
)

// Loading extractor for brand and model of vehicle
var mmExtractor = extractor.NewRegexExtractor("./resources/marque_modele.json")

// QueryMarqueModele is the handler func for "post" request to http://localhost:3000/api/v1/marque-modele, this endpoint can be tested like this:
// curl -X POST http://localhost:3000/api/v1/marque-modele -H "X-API-KEY: 6366d1e468026921bb3194ab" -H "Content-Type: application/json" -d "{\"text\": \"J'ai un Renault Scenic monsieur\"}"
func QueryMarqueModele(c *fiber.Ctx) error {
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
	err = json.Unmarshal(c.Body(), &body)
	if err != nil {
		return err
	}
	// Extracting information or returning empty dict if no one
	if utils.MapContains(body, "text") {
		extractedData = mmExtractor.GetEntity(body["text"])
	}
	res, _ = json.Marshal(extractedData)

	return c.Send(res)
}
