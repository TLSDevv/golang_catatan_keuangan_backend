package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

// get all transactions by user_id (only admin)
func (r Repository) GetTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at, updated_at, deleted_at
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
			&t.UserID,
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
			return nil, err
		}

		trcs = append(trcs, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trcs, nil
}

// get active transactions by user_id (admin, user)
func (r Repository) GetActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at, updated_at, deleted_at
		FROM
			transactions
		WHERE
			user_id = ? AND deleted_at IS NULL`

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
			&t.UserID,
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
			return nil, err
		}

		trcs = append(trcs, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trcs, nil
}

// get inactive transactions by user_id (only admin)
func (r Repository) GetInActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at, updated_at, deleted_at
		FROM
			transactions
		WHERE
			user_id = ? AND deleted_at IS NOT NULL`

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
			&t.UpdatedAt,
			&t.DeletedAt,
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

// get today transactions by user_id (only admin)
func (r Repository) GetTodayTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) {
	sql := `
		SELECT
			id, user_id, transaction_name, category, transaction_type, amount, transaction_at, created_at, updated_at, deleted_at
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
			&t.UpdatedAt,
			&t.DeletedAt,
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
