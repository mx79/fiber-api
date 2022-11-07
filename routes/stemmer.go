package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/controllers"
)

// StemmerRoute
func StemmerRoute(route fiber.Router) {
	route.Post("", controllers.QueryStemmer)
}
