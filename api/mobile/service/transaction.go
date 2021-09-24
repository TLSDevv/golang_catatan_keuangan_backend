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
	DB       *sql.DB
	Validate *validator.Validate
}

var (
	//Register Repo
	transactionRepo = repository.NewTransactionRepository()
)

func NewTransactioService(db *sql.DB, v *validator.Validate) TransactionServiceInterface {
	return &TransactionService{
		DB:       db,
		Validate: v,
	}
}

type TransactionServiceInterface interface {
	ListTransaction(ctx context.Context, int, page int, userId int) []web.TransactionResponse
	GetTransaction(ctx context.Context, idTransaction int) web.TransactionResponse
	CreateTransaction(ctx context.Context, t web.TransactionCreateRequest)
	UpdateTransaction(ctx context.Context, t web.TransactionUpdateRequest)
	DeleteTransaction(ctx context.Context, idTransaction int)
}

func (t *TransactionService) ListTransaction(ctx context.Context, limit int, page int, userId int) []web.TransactionResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transactions, err := transactionRepo.List(ctx, tx, limit, page, userId)
	helper.PanicIfError(err)

	return helper.ToTransactionResponses(transactions)
}

func (t *TransactionService) GetTransaction(ctx context.Context, idTransaction int) web.TransactionResponse {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	transaction, err := transactionRepo.GetByID(ctx, tx, idTransaction)
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

	err = transactionRepo.Store(ctx, tx, transaction)
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

	_, err = transactionRepo.GetByID(ctx, tx, int(transactionRequest.Id))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	transaction := helper.ToTransactionUpdate(transactionRequest)

	err = transactionRepo.Update(ctx, tx, transaction)
	helper.PanicIfError(err)
}

func (t *TransactionService) DeleteTransaction(ctx context.Context, transactionId int) {
	tx, err := t.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	_, err = transactionRepo.GetByID(ctx, tx, transactionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = transactionRepo.Delete(ctx, tx, transactionId)
	helper.PanicIfError(err)
}
