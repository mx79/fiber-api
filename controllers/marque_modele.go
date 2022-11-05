package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/go-nlp/extractor"
	"github.com/mx79/go-nlp/utils"
)

// Loading extractor for brand and model of vehicle
var mmExtractor = extractor.NewRegexExtractor("./resources/marque_modele.json")

// QueryMarqueModele
// curl -X POST http://localhost:3000/api/v1/marque-modele -H "Content-Type: application/json" -d "{\"text\": \"J'ai un Renault Scenic monsieur\"}"
func QueryMarqueModele(c *fiber.Ctx) error {
	var (
		body          map[string]string
		extractedData map[string]interface{}
		res           []byte
	)
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		return err
	}
	if utils.MapContains(body, "text") {
		extractedData = mmExtractor.GetEntity(body["text"])
	}
	res, _ = json.Marshal(extractedData)
	return c.Send(res)
}
