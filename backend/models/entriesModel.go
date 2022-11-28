package models

import (
	"time"
)

type Entry struct {
	AccountNumber string    `json:"account_number"`
	EntryType     bool      `json:"entry_type"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
