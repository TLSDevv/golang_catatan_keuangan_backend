package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

// get transaction by id (only admin)
func (r Repository) GetByID(ctx context.Context, transactionID int) (entities.Transaction, error) {
	sql := `
		SELECT
			id, transaction_name, category, transaction_type, amount, transaction_at, created_at, updated_at, deleted_at
		FROM
			transactions
		WHERE
			id=?`

	t := entities.Transaction{}

	err := r.DB.QueryRowContext(ctx, sql, transactionID).Scan(
		&t.ID,
		&t.TransactionName,
		&t.Category,
		&t.TransactionType,
		&t.Amount,
		&t.TransactionAt,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.DeletedAt,
	)

	if err != nil {
		return entities.Transaction{}, err
	}

	return t, nil
}

// get transaction by id but only active (admin, user)
func (r Repository) GetActiveByID(ctx context.Context, transactionID int) (entities.Transaction, error) {
	sql := `
		SELECT
			id, transaction_name, category, transaction_type, amount, transaction_at, created_at, updated_at, deleted_at
		FROM
			transactions
		WHERE
			id=? AND deleted_at IS NULL`

	t := entities.Transaction{}

	err := r.DB.QueryRowContext(ctx, sql, transactionID).Scan(
		&t.ID,
		&t.TransactionName,
		&t.Category,
		&t.TransactionType,
		&t.Amount,
		&t.TransactionAt,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.DeletedAt,
	)

	if err != nil {
		return entities.Transaction{}, err
	}

	return t, nil
}
