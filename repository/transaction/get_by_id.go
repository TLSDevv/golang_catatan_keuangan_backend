package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) GetByID(ctx context.Context, transactionID int) (*entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at
		FROM
			transactions
		WHERE
			id=$1`

	trc := entities.Transaction{}
	err := r.DB.QueryRowContext(ctx, sql, transactionID).Scan(
		&trc.ID,
		&trc.TransactionName,
		&trc.Category,
		&trc.TransactionType,
		&trc.TransactionAt,
		&trc.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &trc, nil
}
