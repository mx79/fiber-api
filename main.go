package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	env "github.com/joho/godotenv"
	"github.com/mx79/go-nlp/extractor"
	"github.com/mx79/go-nlp/utils"
	"log"
	"os"
)

// Loading extractor
var ext = extractor.NewDefaultRegexExtractor()

func init() {
	// Loading env file
	err := env.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}
}

func main() {
	app := fiber.New()

	// curl -X POST http://localhost:3000/v1/extractor -H "Content-Type: application/json" -d "{\"text\": \"yo on est monday et ici c'est Paris\"}"
	app.Post("/v1/extractor", func(c *fiber.Ctx) error {
		fmt.Println(c.Request())
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
			extractedData = ext.GetEntity(body["text"])
		}
		res, err = json.Marshal(extractedData)
		return c.Send(res)
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
