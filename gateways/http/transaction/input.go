package transaction

import (
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type (
	ResponseBody struct {
		ID              int       `json:"id"`
		TransactionName string    `json:"transaction_name"`
		Category        string    `json:"category"`
		TransactionType int       `json:"transaction_type"`
		Amount          int       `json:"amount"`
		TransactionAt   time.Time `json:"transaction_at"`
		CreatedAt       time.Time `json:"created_at"`
	}

	TransactionListResponse struct {
		Success bool           `json:"success"`
		Data    []ResponseBody `json:"data"`
	}
)

func formatSliceResponse(transactions []entities.Transaction) []ResponseBody {
	transactionResponse := make([]ResponseBody, len(transactions))

	for i, t := range transactions {
		transactionResponse[i] = ResponseBody{
			ID:              t.ID,
			TransactionName: t.TransactionName,
			Category:        t.Category,
			TransactionType: t.TransactionType,
			Amount:          t.Amount,
			TransactionAt:   t.TransactionAt,
			CreatedAt:       t.CreatedAt,
		}
	}

	return transactionResponse
}
