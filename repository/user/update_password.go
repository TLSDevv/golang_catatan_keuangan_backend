package user

import (
	"context"
	"time"
)

func (u UserRepository) UpdatePassword(ctx context.Context, id int, password string) error {
	sql := `
		UPDATE
			users
		SET
			password=?,
			updated_at=?
		WHERE
			id=?`

	updateAt := time.Now()

	_, err := u.DB.ExecContext(ctx, sql,
		password,
		updateAt,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
