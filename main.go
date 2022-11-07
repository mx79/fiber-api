package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	env "github.com/joho/godotenv"
	"github.com/mx79/fiber-api/config"
	"github.com/mx79/fiber-api/routes"
	"log"
	"os"
)

// setupRoutes
// Query examples for each api route
// curl -X POST http://localhost:3000/api/v1/marque-modele -H "Content-Type: application/json" -H "X-API-KEY: 3c05efa8-14c3-4413-834e-e2dfc982fbf9" -d "{\"text\": \"Ouais c'est un renault Mégane de Clio en fait une peugeot ou bien une citroen je sais plus\"}"
// curl -X POST http://localhost:3000/api/v1/stemmer -H "Content-Type: application/json" -H "X-API-KEY: 3c05efa8-14c3-4413-834e-e2dfc982fbf9" -d "{\"text\": \"je me suis trompé de carburant pour ma voiture c'est la catastrophe elle fait un bruit bizarre\"}"
// curl -X POST http://localhost:3000/api/v1/stopword -H "Content-Type: application/json" -H "X-API-KEY: 3c05efa8-14c3-4413-834e-e2dfc982fbf9" -d "{\"text\": \"Là je suis en vacances avec mes parents, j'espère que le programme n'est pas trop long à tourner quand même\"}"
// curl -X POST http://localhost:3000/api/v1/rake -H "Content-Type: application/json" -H "X-API-KEY: 3c05efa8-14c3-4413-834e-e2dfc982fbf9" -d "{\"text\": \"The growing doubt of human autonomy and reason has created a state of moral confusion where man is left without the guidance of either revelation or reason. The result is the acceptance of a relativistic position which proposes that value judgements and ethical norms are exclusively matters of arbitrary preference and that no objectively valid statement can be made in this realm... But since man cannot live without values and norms, this relativism makes him an easy prey for irrational value systems.\"}"
func setupRoutes(app *fiber.App) {
	// API group
	api := app.Group("/api/v1")

	// Adding all routes
	routes.UserRoute(api.Group("/users"))
	routes.MarqueModeleRoute(api.Group("/marque-modele"))
	routes.StemmerRoute(api.Group("/stemmer"))
	routes.StopwordRoute(api.Group("/stopword"))
	routes.RakeRoute(api.Group("/rake"))
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
