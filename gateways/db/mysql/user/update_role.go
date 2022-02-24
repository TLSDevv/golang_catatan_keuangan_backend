package user

import (
	"context"
	"time"
)

func (u UserRepository) UpdateRole(ctx context.Context, id int, role int) error {
	sql := `
		UPDATE
			users
		SET
			role=?,
			updated_at=?
		WHERE
			id=?`

	updateAt := time.Now()

	_, err := u.DB.ExecContext(ctx, sql,
		role,
		updateAt,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
