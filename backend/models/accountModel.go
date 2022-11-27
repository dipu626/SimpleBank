package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id"`
	Email     *string            `json:"email" validate:"required"`
	Name      *string            `json:"name" validate:"required,min=2,max=10"`
	Ballance  *float64           `json:"price" validate:"required,min=0"`
	CreatedAt time.Time          `json:"created_at"`
	PhoneNo   *string            `json:"phone_no"`
}
