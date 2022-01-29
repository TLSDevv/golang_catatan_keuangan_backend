package repository

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/entity"
)

type transactionRepository struct {
}

// type TransactionRepository interface {
// 	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Transaction, error)
// }

func NewTransactionRepository() transactionRepository {
	return transactionRepository{}
}

func (t *transactionRepository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Transaction, error) {
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
