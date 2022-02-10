package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) Create(ctx context.Context, trc entities.Transaction) error {
	sql := `
		INSERT INTO
			transactions(
				user_id,
				transaction_name,
				category,
				transaction_type,
				amount,
				transaction_at
			)
			VALUES($1, $2, $3, $4, $5, $6)`

	_, err := r.DB.ExecContext(ctx, sql,
		trc.UserId,
		trc.TransactionName,
		trc.Category,
		trc.TransactionType,
		trc.Amount,
		trc.TransactionAt)

	if err != nil {
		return err
	}

	return nil
}
