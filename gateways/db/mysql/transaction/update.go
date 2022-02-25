package transaction

import (
	"context"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) Update(ctx context.Context, trc entities.Transaction, tID int) error {
	sql := `
			UPDATE
				transactions
			SET
				user_id=?,
				transaction_name=?,
				category=?,
				transaction_type=?,
				amount=?,
				transaction_at=?
			WHERE
				id=?`

	result, err := r.DB.ExecContext(ctx, sql,
		trc.TransactionName,
		trc.Category,
		trc.TransactionType,
		trc.Amount,
		trc.TransactionAt.Local(),
		tID)

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return errors.New("update failed, no rows affected")
	}

	return nil
}
