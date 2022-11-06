package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Email     string             `json:"email"`
	ApiKey    string             `json:"api_key"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
