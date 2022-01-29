package repository

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/entity"
)

type TransactionRepository struct {
}

type ITransactionRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Transaction, error)
	FindById(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
	FindAllTimeByUserId(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
	FindTodayTransactionByUserId(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
	FindWeekTransactionByUserId(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
	FindMonthTransactionByUserId(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
	Create(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
	Update(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
	Delete(ctx context.Context, tx *sql.Tx, transaction entity.Transaction) error
}

func NewTransactionRepository() TransactionRepository {
	return TransactionRepository{}
}

func (t *TransactionRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Transaction, error) {
	sql := `SELECT 
				ID, TrcName, Category, TrcType, Amount, TransactionAt, CreatedAt
			FROM
				transactions
		`
	rows, err := tx.QueryContext(ctx, sql)
	defer rows.Close()

	if err == nil {
		panic(err)
	}

	transactions := []entity.Transaction{}
	for rows.Next() {
		transaction := entity.Transaction{}

		err := rows.Scan(
			&transaction.ID,
			&transaction.TrcName,
			&transaction.Category,
			&transaction.TrcType,
			&transaction.Amount,
			&transaction.TransactionAt,
			&transaction.CreatedAt,
		)

		if err == nil {
			panic(err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
