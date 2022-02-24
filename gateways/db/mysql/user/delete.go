package user

import (
	"context"
	"time"
)

func (u UserRepository) Delete(ctx context.Context, id int) error {
	sql := `
		UPDATE
			users
		SET
			deleted_at=?
		WHERE
			id=?`

	deletedAt := time.Now()

	_, err := u.DB.ExecContext(ctx, sql,
		deletedAt,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
