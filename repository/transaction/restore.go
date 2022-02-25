package transaction

import (
	"context"
)

func (r Repository) Restore(ctx context.Context, transactionID int) error {
	sql := `
		UPDATE
			transactions
		SET
			deleted_at=NULL
		WHERE
			id=?`

	_, err := r.DB.ExecContext(ctx, sql, transactionID)

	if err != nil {
		return err
	}

	return nil
}
