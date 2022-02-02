package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/transaction"
)

type TransactionService struct {
	TransactionRepository transaction.Repository
	DB                    *sql.DB
}

func NewTransactionService(transactionRepo transaction.Repository, db *sql.DB) TransactionService {
	return TransactionService{
		TransactionRepository: transactionRepo,
		DB:                    db,
	}
}

func (service TransactionService) FindAll(ctx context.Context) ([]entities.Transaction, error) {
	tx, err := service.DB.Begin()

	if err != nil {
		return []entities.Transaction{}, nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	trcList, err := service.TransactionRepository.FindAll(ctx, tx)
	if err != nil {
		return []entities.Transaction{}, nil
	}

	trcDtoList := []entities.Transaction{}
	for _, tItem := range trcList {
		trcDtoList = append(trcDtoList, entities.Transaction{
			ID:            tItem.ID,
			UserId:        tItem.UserId,
			TrcName:       tItem.TrcName,
			Category:      tItem.Category,
			TrcType:       tItem.TrcType,
			TransactionAt: tItem.TransactionAt,
			CreatedAt:     tItem.CreatedAt,
		})
	}

	return trcDtoList, nil
}
