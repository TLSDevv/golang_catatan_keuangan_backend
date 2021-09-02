package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
)

type transactionRepo struct {
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepo{}
}

var structureTransaction string = `id,name_transaction,type_transaction,category_id,amount,description,created_at,updated_at`
var structureTransactionStore string = `id,user_id,name_transaction,type_transaction,category_id,amount,description,created_at,updated_at,deleted_at`
var structureTransactionUpdate string = `name_transaction,type_transaction,category_id,amount,description,updated_at`

func (t *transactionRepo) ListByUser(ctx context.Context, tx *sql.Tx, limit int, page int, userId int) ([]domain.Transaction, error) {
	transactions := []domain.Transaction{}

	var err error
	var rows *sql.Rows

	if limit != 0 {
		if page != 0 {
			offset := (page - 1) * limit
			sql := `SELECT ` + structureTransaction + `FROM transactions WHERE user_id=$3 AND deleted_at IS NOT NULL LIMIT $1 OFFSET $2`
			rows, err = tx.QueryContext(ctx, sql, limit, offset, userId)
		} else {
			sql := `SELECT ` + structureTransaction + `FROM transactions WHERE user_id=$3 deleted_at IS NOT NULL LIMIT $1`
			rows, err = tx.QueryContext(ctx, sql, limit, userId)
		}
	} else {
		sql := `SELECT ` + structureTransaction + `FROM transactions WHERE user_id=$3 deleted_at IS NOT NULL`
		rows, err = tx.QueryContext(ctx, sql, userId)
	}

	helper.PanicIfError(err)

	defer rows.Close()

	for rows.Next() {
		var transaction domain.Transaction
		err = rows.Scan(
			&transaction.Id,
			&transaction.NameTransaction,
			&transaction.TypeTransaction,
			&transaction.CategoryId,
			&transaction.Amount,
			&transaction.Description,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		helper.PanicIfError(err)
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (t *transactionRepo) GetByID(ctx context.Context, tx *sql.Tx, transactionId int) (domain.Transaction, error) {
	sql := `SELECT ` + structureTransaction + ` FROM transactions WHERE id=$1 AND deleted_at IS NULL`
	rows, err := tx.QueryContext(ctx, sql, transactionId)

	helper.PanicIfError(err)

	transaction := domain.Transaction{}
	if rows.Next() {
		err := rows.Scan(
			&transaction.Id,
			&transaction.NameTransaction,
			&transaction.TypeTransaction,
			&transaction.CategoryId,
			&transaction.Amount,
			&transaction.Description,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		helper.PanicIfError(err)

		return transaction, nil
	} else {
		return transaction, errors.New("Transaction Not Found")
	}
}

func (t *transactionRepo) Store(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) error {
	transaction.CreatedAt = time.Now().Local()
	transaction.UpdatedAt = transaction.CreatedAt
	sql := `INSERT INTO transactions (
		` + structureTransactionStore + `)
		values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, err := tx.ExecContext(
		ctx,
		sql,
		transaction.Id,
		transaction.UserId,
		transaction.TypeTransaction,
		transaction.CategoryId,
		transaction.Amount,
		transaction.Description,
		transaction.CreatedAt,
		transaction.UpdatedAt,
		transaction.DeletedAt)

	helper.PanicIfError(err)

	return nil
}

func (t *transactionRepo) Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) error {
	transaction.UpdatedAt = time.Now().Local()
	sql := `UPDATE INTO transactions (
		` + structureTransactionUpdate + `)
		values ($1,$2,$3,$4,$5,$6) WHERE id=$8`

	_, err := tx.ExecContext(
		ctx,
		sql,
		transaction.NameTransaction,
		transaction.TypeTransaction,
		transaction.CategoryId,
		transaction.Amount,
		transaction.Description,
		transaction.UpdatedAt,
		transaction.Id)

	helper.PanicIfError(err)

	return nil
}

func (t *transactionRepo) Delete(ctx context.Context, tx *sql.Tx, transactionId int) error {
	sql := `UPDATE INTO transactions (deleted_at)
		values ($1) WHERE id=$2`

	deleteAt := time.Now().Local()

	_, err := tx.ExecContext(
		ctx,
		sql,
		deleteAt,
		transactionId)

	helper.PanicIfError(err)

	return nil
}
