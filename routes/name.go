package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/controllers"
)

func NameRoute(route fiber.Router) {
	route.Get("", controllers.GetNames)
}
