package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/controllers"
)

// StopwordRoute
func StopwordRoute(route fiber.Router) {
	route.Post("", controllers.QueryStopword)
}
