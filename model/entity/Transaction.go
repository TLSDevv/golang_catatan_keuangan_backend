package entity

import (
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/dto"
)

type Transaction struct {
	ID            int       `json:"id"`
	UserId        string    `json:"user_id"`
	TrcName       string    `json:"trc_name"`
	Category      string    `json:"category"`
	TrcType       int       `json:"trc_type"`
	Amount        int       `json:"amount"`
	TransactionAt time.Time `json:"transaction_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

func (t Transaction) ToDTO() dto.Transaction {
	return dto.Transaction{
		ID:            t.ID,
		TrcName:       t.TrcName,
		Category:      t.Category,
		TrcType:       t.TrcType,
		Amount:        t.Amount,
		TransactionAt: t.TransactionAt,
		CreatedAt:     t.CreatedAt,
	}
}
