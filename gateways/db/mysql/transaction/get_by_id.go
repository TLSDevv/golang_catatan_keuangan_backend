package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) GetByID(ctx context.Context, transactionID int) (entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at
		FROM
			transactions
		WHERE
			id=?`

	t := entities.Transaction{}

	err := r.DB.QueryRowContext(ctx, sql, transactionID).Scan(
		&t.ID,
		&t.UserID,
		&t.TransactionName,
		&t.Category,
		&t.TransactionType,
		&t.Amount,
		&t.TransactionAt,
		&t.CreatedAt,
	)

	if err != nil {
		return entities.Transaction{}, err
	}

	return t, nil
}
