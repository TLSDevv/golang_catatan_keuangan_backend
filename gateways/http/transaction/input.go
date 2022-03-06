package transaction

import (
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type (
	TransactionRequest struct {
		TransactionName string `json:"transaction_name"`
		Category        string `json:"category"`
		TransactionType int    `json:"transaction_type"`
		Amount          int    `json:"amount"`
		TransactionAt   string `json:"transaction_at"`
	}

	TransactionRequestValidationError struct {
		TransactionName string `json:"transaction_name,omitempty"`
		Category        string `json:"category,omitempty"`
		TransactionType int    `json:"transaction_type,omitempty"`
		Amount          int    `json:"amount,omitempty"`
		TransactionAt   string `json:"transaction_at,omitempty"`
	}

	ResponseBody struct {
		ID              int               `json:"id"`
		UserID          int               `json:"user_id"`
		TransactionName string            `json:"transaction_name"`
		Category        string            `json:"category"`
		TransactionType int               `json:"transaction_type"`
		Amount          int               `json:"amount"`
		TransactionAt   time.Time         `json:"transaction_at"`
		CreatedAt       time.Time         `json:"created_at"`
		UpdatedAt       time.Time         `json:"updated_at"`
		DeletedAt       entities.NullTime `json:"deleted_at"`
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
			UserID:          t.UserID,
			TransactionName: t.TransactionName,
			Category:        t.Category,
			TransactionType: t.TransactionType,
			Amount:          t.Amount,
			TransactionAt:   t.TransactionAt,
			CreatedAt:       t.CreatedAt,
			UpdatedAt:       t.UpdatedAt,
			DeletedAt:       t.DeletedAt,
		}
	}

	return transactionResponse
}
