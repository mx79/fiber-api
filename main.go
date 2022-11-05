package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	env "github.com/joho/godotenv"
	"github.com/mx79/fiber-api/config"
	"github.com/mx79/go-nlp/extractor"
	"github.com/mx79/go-nlp/utils"
	"log"
	"os"
)

// Loading extractor
var mmExtractor = extractor.NewRegexExtractor("./resources/marque_modele.json")

func setupRoutes(app *fiber.App) {
	// curl -X POST http://localhost:3000/v1/extractor -H "Content-Type: application/json" -d "{\"text\": \"yo on est monday et ici c'est Paris\"}"
	app.Post("/v1/name", func(c *fiber.Ctx) error {
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
			extractedData = mmExtractor.GetEntity(body["text"])
		}
		res, _ = json.Marshal(extractedData)
		return c.Send(res)
	})
}

func main() {
	// Init app
	app := fiber.New()
	app.Use(logger.New())

	// Loading env file
	err := env.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}

	// Connecting to MongoDB cluster
	config.ConnectDB()

	// Setup routes
	setupRoutes(app)

	// Starting server on port specified and catch error if any
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
