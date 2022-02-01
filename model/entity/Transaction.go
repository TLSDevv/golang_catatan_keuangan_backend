package entity

import (
	"errors"
	"time"
)

type Transaction struct {
	ID            int       `json:"id"`
	UserId        int       `json:"user_id"`
	TrcName       string    `json:"trc_name"`
	Category      string    `json:"category"`
	TrcType       int       `json:"trc_type"`
	Amount        int       `json:"amount"`
	TransactionAt time.Time `json:"transaction_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func (t Transaction) Validate() error {
	if t.UserId == 0 {
		return errors.New("user_id is required")
	}

	if len(t.TrcName) == 0 {
		return errors.New("tc_name is required")
	}

	if len(t.Category) == 0 {
		return errors.New("category is required")
	}

	if t.TrcType == 0 {
		return errors.New("trc_type is required")
	}

	if t.Amount == 0 {
		return errors.New("amount is required")
	}

	return nil
}

// func (t Transaction) ToDTO() dto.Transaction {
// 	return dto.Transaction{
// 		ID:            t.ID,
// 		TrcName:       t.TrcName,
// 		Category:      t.Category,
// 		TrcType:       t.TrcType,
// 		Amount:        t.Amount,
// 		TransactionAt: t.TransactionAt,
// 		CreatedAt:     t.CreatedAt,
// 	}
// }

func (t Transaction) Update() {
	t.Validate()

	t.UpdatedAt = time.Now()
}

func (t Transaction) Delete() {
	t.DeletedAt = time.Now()
}
