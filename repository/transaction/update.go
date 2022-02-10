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
				trc_name=$1,
				category=$2,
				trc_type=$3,
				amount=$4,
				transaction_at=$5
			WHERE
				id=$6`

	_, err := r.DB.ExecContext(ctx, sql,
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
