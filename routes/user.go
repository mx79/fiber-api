package routes

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO: Create route handler here

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Email     string             `json:"email"`
}
