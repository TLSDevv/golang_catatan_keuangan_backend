package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) Update(ctx context.Context, trc entities.Transaction) error {
	sql := `
			UPDATE
				transactions
			SET
				user_id=?
				trc_name=?,
				category=?,
				trc_type=?,
				amount=?,
				transaction_at=?
			WHERE
				id=?`

	_, err := r.DB.ExecContext(ctx, sql,
		trc.UserID,
		trc.TransactionName,
		trc.Category,
		trc.TransactionType,
		trc.Amount,
		trc.TransactionAt,
		trc.ID)

	if err != nil {
		return err
	}

	return nil
}
