package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type TransactionService struct {
	TransactionRepo repository.TransactionRepository
	DB              *sql.DB
}

func NewService(transactionRepository repository.TransactionRepository, db *sql.DB) TransactionServiceInterface {
	return &TransactionService{
		TransactionRepo: transactionRepository,
		DB:              db,
	}
}

func (t *TransactionService) ListTransaction(ctx context.Context, limit int, page int) []web.TransactionResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transactions, err := t.TransactionRepo.List(ctx, tx, limit, page)
	helper.PanicIfError(err)

	return helper.ToTransactionResponses(transactions)
}

func (t *TransactionService) GetTransaction(ctx context.Context, idTransaction int) web.TransactionResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transaction, err := t.TransactionRepo.GetByID(ctx, tx, idTransaction)
	helper.PanicIfError(err)

	return helper.ToTransactionResponse(transaction)
}

func (t *TransactionService) CreateTransaction(ctx context.Context, transactionRequest web.TransactionCreateRequest) {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transaction := helper.ToTransaction(transactionRequest)

	err = t.TransactionRepo.Store(ctx, tx, transaction)
	helper.PanicIfError(err)
}

func (t *TransactionService) UpdateTransaction(ctx context.Context, idTransaction int, transactionRequest web.TransactionCreateRequest) {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transaction := helper.ToTransaction(transactionRequest)

	err = t.TransactionRepo.Update(ctx, tx, idTransaction, transaction)
	helper.PanicIfError(err)
}

func (t *TransactionService) DeleteTransaction(ctx context.Context, idTransaction int) {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	err = t.TransactionRepo.Delete(ctx, tx, idTransaction)
	helper.PanicIfError(err)
}
