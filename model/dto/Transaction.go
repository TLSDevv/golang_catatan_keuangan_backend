package dto

import "time"

type Transaction struct {
	ID            int       `json:"id"`
	UserId        int       `json:"user_id"`
	TrcName       string    `json:"trc_name"`
	Category      string    `json:"category"`
	TrcType       int       `json:"trc_type"`
	Amount        int       `json:"amount"`
	TransactionAt time.Time `json:"transaction_at"`
	CreatedAt     time.Time `json:"created_at"`
}
