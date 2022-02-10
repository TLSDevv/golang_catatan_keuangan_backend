package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) GetTodayTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at
		FROM
			transactions
		WHERE
			user_id = ? AND
			(transaction_at >= ? AND transaction_at <= ?)`

	year, month, day := time.Now().Date()

	today_start := fmt.Sprintf("%d-%s-%d", year, month, day)
	today_end := fmt.Sprintf("%s %s", today_start, "23:59:59")

	rows, err := r.DB.QueryContext(ctx, sql, userID, today_start, today_end)
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
