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

func setupRoutes(app *fiber.App) {
	// API group
	api := app.Group("/api/v1")

	// Adding all routes
	routes.UserRoute(api.Group("/users"))
	routes.MarqueModeleRoute(api.Group("/marque-modele"))
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
