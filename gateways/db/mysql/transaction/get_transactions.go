package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) GetTransactions(ctx context.Context) ([]entities.Transaction, error) {
	sql := `
		SELECT
			id, transaction_name, category, transaction_type, amount, transaction_at, created_at
		FROM
			transactions`

	rows, err := r.DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := []entities.Transaction{}

	for rows.Next() {
		t := entities.Transaction{}

		err := rows.Scan(
			&t.ID,
			&t.TransactionName,
			&t.Category,
			&t.TransactionType,
			&t.Amount,
			&t.TransactionAt,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
