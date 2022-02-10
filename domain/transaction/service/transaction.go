package service

import (
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
)

var _ transaction.Service = Transaction{}

type Transaction struct {
	repo transaction.Repository
}

func NewTransactionService(repo transaction.Repository, db *sql.DB) *Transaction {
	return &Transaction{
		repo: repo,
	}
}
