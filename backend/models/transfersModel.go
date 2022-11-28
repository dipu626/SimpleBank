package models

import "time"

type Transfer struct {
	AccountFrom string    `json:"account_from"`
	AccountTo   string    `json:"account_to"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}
