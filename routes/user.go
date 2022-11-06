package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/controllers"
)

// UserRoute
func UserRoute(route fiber.Router) {
	//route.Get("/:id")                      // Get user by id
	route.Post("", controllers.CreateUser)       // Create user
	route.Put("/:id", controllers.UpdateUser)    // Update user
	route.Delete("/:id", controllers.DeleteTodo) // Delete user
}
