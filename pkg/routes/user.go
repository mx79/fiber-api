package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/pkg/config"
	"github.com/mx79/fiber-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

// UserRoute set up api route for users management
func UserRoute(route fiber.Router) {
	//route.Get("/:id")                      // Get user by id
	route.Post("", createUser)       // Create user
	route.Put("/:id", updateUser)    // Update user
	route.Delete("/:id", deleteUser) // Delete user
}

// createUser accept a POST request to create a new user in database
//
//	curl -X POST http://localhost:3000/api/v1/users
//	-H "Content-Type: application/json"
//	-d "{\"first_name\": \"Max\", \"last_name\": \"Lesage\", \"email\": \"max@test.fr\", \"password\": \"test\"}"
func createUser(c *fiber.Ctx) error {
	// Get users collection and parse the body of the request
	users := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))
	data := new(models.User)
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// Verify if first_name, last_name, email and password are in the body
	if data.FirstName == "" || data.LastName == "" || data.Email == "" || data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Some required keys are not filed or not in request body",
			"error":   err,
		})
	}

	// Verify now if entered email is already in database
	emailTest := &models.User{}
	err = users.FindOne(c.Context(), bson.M{"email": data.Email}).Decode(emailTest)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "The entered email adress already exists",
			"error":   err,
		})
	}

	// Set user attributes
	data.ID = primitive.NewObjectID()
	data.Password, _ = hashPassword(data.Password)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	// Try to insert the user in database
	result, err := users.InsertOne(c.Context(), data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert user",
			"error":   err,
		})
	}

	// Get the inserted data
	user := &models.User{}
	_ = users.FindOne(c.Context(), bson.M{"_id": result.InsertedID}).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}

//// GetUser shows information of the user with the request id
// curl -X GET http://localhost:3000/api/v1/users/636798a20343d30036d343ec
//func GetTodo(c *fiber.Ctx) error {
//	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
//
//	// get parameter value
//	paramID := c.Params("id")
//
//	// convert parameterID to objectId
//	id, err := primitive.ObjectIDFromHex(paramID)
//
//	// if error while parsing paramID
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Cannot parse Id",
//			"error":   err,
//		})
//	}
//
//	// find todo and return
//
//	todo := &models.Todo{}
//
//	query := bson.D{{Key: "_id", Value: id}}
//
//	err = todoCollection.FindOne(c.Context(), query).Decode(todo)
//
//	if err != nil {
//		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//			"success": false,
//			"message": "Todo Not found",
//			"error":   err,
//		})
//	}
//
//	return c.Status(fiber.StatusOK).JSON(fiber.Map{
//		"success": true,
//		"data": fiber.Map{
//			"todo": todo,
//		},
//	})
//}

// updateUser updates the requested user attributes in database
//
//	curl -X PUT http://localhost:3000/api/v1/users/636798a20343d30036d343ec
//	-H "Content-Type: application/json"
//	-d "{\"first_name\": \"Jean\", \"last_name\": \"Troll\"}"
func updateUser(c *fiber.Ctx) error {
	// Get users collection and parse parameters
	users := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))
	paramID := c.Params("id")

	// Convert parameter to object id
	id, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// Request content
	data := new(models.User)
	err = c.BodyParser(&data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	// Data to updates on user profile
	var dataToUpdate bson.D
	if data.FirstName != "" {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "first_name", Value: data.FirstName})
	}
	if data.LastName != "" {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "last_name", Value: data.LastName})
	}
	dataToUpdate = append(dataToUpdate, bson.E{Key: "updated_at", Value: time.Now()})
	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	// Updating user attributes
	err = users.FindOneAndUpdate(c.Context(), bson.M{"_id": id}, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User Not found",
				"error":   err,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update this user",
			"error":   err,
		})
	}

	// Get updated data
	user := &models.User{}
	_ = users.FindOne(c.Context(), bson.M{"_id": id}).Decode(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"user": user,
		},
	})
}

// deleteUser deletes the user with the corresponding id from database
//
//	curl -X DELETE http://localhost:3000/api/v1/users/636798a20343d30036d343ec
func deleteUser(c *fiber.Ctx) error {
	// Get users collection and parse parameters
	users := config.MI.DB.Collection(os.Getenv("USER_COLLECTION"))
	paramID := c.Params("id")

	// Convert parameter to object id
	id, err := primitive.ObjectIDFromHex(paramID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	// Find and delete user
	err = users.FindOneAndDelete(c.Context(), bson.M{"_id": id}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "User Not found",
				"error":   err,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete this user",
			"error":   err,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
