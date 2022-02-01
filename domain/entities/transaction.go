package entities

import (
	"errors"
	"time"
)

var (
	ErrUIdRequired      = errors.New("User id is required")
	ErrTNameRequired    = errors.New("Transaction name is required")
	ErrCategoryRequired = errors.New("Category is required")
	ErrTTypeRequired    = errors.New("Transaction type is required")
	ErrAmountRequired   = errors.New("Amount is required")

	TransactionDomainErrors = []error{
		ErrUIdRequired,
		ErrTNameRequired,
		ErrCategoryRequired,
		ErrTTypeRequired,
		ErrAmountRequired,
	}
)

// entity
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
		return ErrUIdRequired
	}

	if len(t.TrcName) == 0 {
		return ErrTNameRequired
	}

	if len(t.Category) == 0 {
		return ErrCategoryRequired
	}

	if t.TrcType == 0 {
		return ErrTTypeRequired
	}

	if t.Amount == 0 {
		return ErrAmountRequired
	}

	return nil
}

// response data
type TransactionResponse struct {
	ID            int       `json:"id"`
	UserId        int       `json:"user_id"`
	TrcName       string    `json:"trc_name"`
	Category      string    `json:"category"`
	TrcType       int       `json:"trc_type"`
	Amount        int       `json:"amount"`
	TransactionAt time.Time `json:"transaction_at"`
	CreatedAt     time.Time `json:"created_at"`
}

// request input
type TransactionInput struct {
	UserId        int       `json:"user_id"`
	TrcName       string    `json:"trc_name"`
	Category      string    `json:"category"`
	TrcType       int       `json:"trc_type"`
	Amount        int       `json:"amount"`
	TransactionAt time.Time `json:"transaction_at"`
	CreatedAt     time.Time `json:"created_at"`
}

func (t Transaction) Update() {
	t.Validate()

	t.UpdatedAt = time.Now()
}

func (t Transaction) Delete() {
	t.DeletedAt = time.Now()
}
