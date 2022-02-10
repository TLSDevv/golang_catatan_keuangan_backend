package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) GetTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at
		FROM
			transactions
		WHERE
			user_id = ?`

	rows, err := r.DB.QueryContext(ctx, sql, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	trcs := []entities.Transaction{}

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

		trcs = append(trcs, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trcs, nil
}
