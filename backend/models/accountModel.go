package models

import (
	"time"
)

type Account struct {
	AccountNumber string    `json:"account_number"`
	Email         string    `json:"email" validate:"required"`
	Name          string    `json:"name" validate:"required,min=2,max=10"`
	Ballance      float64   `json:"price" validate:"required,min=0"`
	CreatedAt     time.Time `json:"created_at"`
	PhoneNo       string    `json:"phone_no"`
}
