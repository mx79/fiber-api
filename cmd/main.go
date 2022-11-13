package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mx79/fiber-api/pkg/config"
	routes2 "github.com/mx79/fiber-api/pkg/routes"
	"log"
	"os"
)

// setupRoutes set up all api route declared in the package "routes"
func setupRoutes(app *fiber.App) {
	// API group
	api := app.Group("/api/v1")

	// Adding all routes
	routes2.UserRoute(api.Group("/users"))
	routes2.StemmerRoute(api.Group("/stemmer"))
	routes2.StopwordRoute(api.Group("/stopword"))
	routes2.RakeRoute(api.Group("/rake"))
	routes2.SegmenterRoute(api.Group("/segmenter"))
	routes2.PosTaggingRoute(api.Group("/pos-tagging"))
	routes2.NerRoute(api.Group("/ner"))
	routes2.SentimentRoute(api.Group("/sentiment"))
	routes2.WerRoute(api.Group("/wer"))
}

func main() {
	// Init app
	app := fiber.New()
	app.Use(logger.New())

	// Loading env file
	//err := env.Load(".env")
	//if err != nil {
	//	log.Fatal("Unable to load .env file")
	//}

	// Connecting to MongoDB cluster
	config.ConnectDB()

	// Setup routes
	setupRoutes(app)

	// Starting server on port specified and catch error if any
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
