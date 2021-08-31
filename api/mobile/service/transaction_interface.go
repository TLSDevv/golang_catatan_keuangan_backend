package service

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type TransactionServiceInterface interface {
	ListTransaction(limit int, page int) ([]web.TransactionResponse, error)
	GetTransaction(id int) (web.TransactionResponse, error)
	CreateTransaction(t *web.TransactionCreateRequest) error
	UpdateTransaction(t *web.TransactionUpdateRequest) error
	DeleteTransaction(id int) error
}
