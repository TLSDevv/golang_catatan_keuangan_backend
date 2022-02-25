package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (r Repository) Create(ctx context.Context, t entities.Transaction) error {
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
			VALUES(?, ?, ?, ?, ?, ?);
		`
	_, err := r.DB.ExecContext(ctx, sql,
		ctx.Value("user_id").(int),
		t.TransactionName,
		t.Category,
		t.TransactionType,
		t.Amount,
		t.TransactionAt)

	if err != nil {
		return err
	}

	return nil
}
