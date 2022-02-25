package transaction

import (
	"context"
)

func (r Repository) Purge(ctx context.Context, transactionID int) error {
	sql := `
		DELETE FROM
			transactions
		WHERE
			id=?`

	_, err := r.DB.ExecContext(ctx, sql, transactionID)

	if err != nil {
		return err
	}

	return nil
}
