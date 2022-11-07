package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/controllers"
)

// RakeRoute
func RakeRoute(route fiber.Router) {
	route.Post("", controllers.QueryRake)
}
