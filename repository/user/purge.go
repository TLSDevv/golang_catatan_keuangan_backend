package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (u UserRepository) Purge(ctx context.Context, user entities.User) error {
	sql := `
		DELETE FROM
			users
		WHERE
			id=?`

	_, err := u.DB.ExecContext(ctx, sql,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}
