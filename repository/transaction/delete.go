package transaction

import (
	"context"
	"time"
)

func (r Repository) Delete(ctx context.Context, transactionID int) error {
	sql := `
		UPDATE
			transactions
		SET
			deleted_at=?
		WHERE
			id=?`

	_, err := r.DB.ExecContext(ctx, sql,
		time.Now(),
		transactionID)

	if err != nil {
		return err
	}

	return nil
}
