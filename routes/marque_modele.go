package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/controllers"
)

// MarqueModeleRoute
func MarqueModeleRoute(route fiber.Router) {
	route.Post("", controllers.QueryMarqueModele)
}
