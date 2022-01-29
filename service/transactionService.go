package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/dto"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/repository"
)

type ITransactionService interface {
	FindAll(ctx context.Context) ([]dto.Transaction, error)
}

type TransactionService struct {
	TransactionRepository repository.ITransactionRepository
	DB                    *sql.DB
}

func NewTransactionService(transactionRepo repository.ITransactionRepository, db *sql.DB) ITransactionService {
	return TransactionService{
		TransactionRepository: transactionRepo,
		DB:                    db,
	}
}

func (service TransactionService) FindAll(ctx context.Context) ([]dto.Transaction, error) {
	tx, err := service.DB.Begin()

	if err != nil {
		return []dto.Transaction{}, nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	transactions, err := service.TransactionRepository.FindAll(ctx, tx)

	return helper.TransactionsToDTO(transactions), nil
}
