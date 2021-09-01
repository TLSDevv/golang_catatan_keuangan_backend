package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/exception"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
	"github.com/go-playground/validator"
)

type TransactionService struct {
	TransactionRepo repository.TransactionRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewTransactioService(transactionRepository repository.TransactionRepository, db *sql.DB, v *validator.Validate) TransactionServiceInterface {
	return &TransactionService{
		TransactionRepo: transactionRepository,
		DB:              db,
		Validate:        v,
	}
}

func (t *TransactionService) ListTransaction(ctx context.Context, limit int, page int, userId int) []web.TransactionResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transactions, err := t.TransactionRepo.ListByUser(ctx, tx, limit, page, userId)
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
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTransactionResponse(transaction)
}

func (t *TransactionService) CreateTransaction(ctx context.Context, transactionRequest web.TransactionCreateRequest) {
	err := t.Validate.Struct(transactionRequest)
	helper.PanicIfError(err)

	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transaction := helper.ToTransactionCreate(transactionRequest)

	err = t.TransactionRepo.Store(ctx, tx, transaction)
	helper.PanicIfError(err)
}

func (t *TransactionService) UpdateTransaction(ctx context.Context, transactionRequest web.TransactionUpdateRequest) {
	err := t.Validate.Struct(transactionRequest)
	helper.PanicIfError(err)

	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	_, err = t.TransactionRepo.GetByID(ctx, tx, int(transactionRequest.Id))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	transaction := helper.ToTransactionUpdate(transactionRequest)

	err = t.TransactionRepo.Update(ctx, tx, transaction)
	helper.PanicIfError(err)
}

func (t *TransactionService) DeleteTransaction(ctx context.Context, transactionId int) {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	_, err = t.TransactionRepo.GetByID(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = t.TransactionRepo.Delete(ctx, tx, transactionId)
	helper.PanicIfError(err)
}
