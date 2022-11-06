package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mx79/fiber-api/config"
	"github.com/mx79/fiber-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
)

// CreateUser accept a POST request to create a new user in database
// curl -X POST http://localhost:3000/api/v1/marque-modele
// -H "Content-Type: application/json"
// -d "{\"text\": \"J'ai un Renault Scenic monsieur\"}"
func CreateUser(c *fiber.Ctx) error {
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
	// Set user attributes
	data.ID = primitive.NewObjectID()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
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
	query := bson.M{"_id": result.InsertedID}
	users.FindOne(c.Context(), query).Decode(user)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": user,
		},
	})
}

//// GetTodo : get a single todo
//// PARAM: id
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
//
//// UpdateTodo : Update a todo
//// PARAM: id
//func UpdateTodo(c *fiber.Ctx) error {
//	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
//
//	// find parameter
//	paramID := c.Params("id")
//
//	// convert parameterID to objectId
//	id, err := primitive.ObjectIDFromHex(paramID)
//
//	// if parameter cannot parse
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Cannot parse id",
//			"error":   err,
//		})
//	}
//
//	// var data Request
//	data := new(models.Todo)
//	err = c.BodyParser(&data)
//
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Cannot parse JSON",
//			"error":   err,
//		})
//	}
//
//	query := bson.D{{Key: "_id", Value: id}}
//
//	// updateData
//	var dataToUpdate bson.D
//
//	if data.Title != nil {
//		// todo.Title = *data.Title
//		dataToUpdate = append(dataToUpdate, bson.E{Key: "title", Value: data.Title})
//	}
//
//	if data.Completed != nil {
//		// todo.Completed = *data.Completed
//		dataToUpdate = append(dataToUpdate, bson.E{Key: "completed", Value: data.Completed})
//	}
//
//	dataToUpdate = append(dataToUpdate, bson.E{Key: "updatedAt", Value: time.Now()})
//
//	update := bson.D{
//		{Key: "$set", Value: dataToUpdate},
//	}
//
//	// update
//	err = todoCollection.FindOneAndUpdate(c.Context(), query, update).Err()
//
//	if err != nil {
//		if err == mongo.ErrNoDocuments {
//			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//				"success": false,
//				"message": "Todo Not found",
//				"error":   err,
//			})
//		}
//
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Cannot update todo",
//			"error":   err,
//		})
//	}
//
//	// get updated data
//	todo := &models.Todo{}
//
//	todoCollection.FindOne(c.Context(), query).Decode(todo)
//
//	return c.Status(fiber.StatusOK).JSON(fiber.Map{
//		"success": true,
//		"data": fiber.Map{
//			"todo": todo,
//		},
//	})
//}
//
//// DeleteTodo : Delete a todo
//// PARAM: id
//func DeleteTodo(c *fiber.Ctx) error {
//	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))
//
//	// get param
//	paramID := c.Params("id")
//
//	// convert parameter to object id
//	id, err := primitive.ObjectIDFromHex(paramID)
//
//	// if parameter cannot parse
//	if err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Cannot parse id",
//			"error":   err,
//		})
//	}
//
//	// find and delete todo
//	query := bson.D{{Key: "_id", Value: id}}
//
//	err = todoCollection.FindOneAndDelete(c.Context(), query).Err()
//
//	if err != nil {
//		if err == mongo.ErrNoDocuments {
//			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
//				"success": false,
//				"message": "Todo Not found",
//				"error":   err,
//			})
//		}
//
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"success": false,
//			"message": "Cannot delete todo",
//			"error":   err,
//		})
//	}
//
//	return c.SendStatus(fiber.StatusNoContent)
//}
