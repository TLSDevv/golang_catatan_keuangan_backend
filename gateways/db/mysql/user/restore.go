package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (u UserRepository) Restore(ctx context.Context, user entities.User) error {
	sql := `
		UPDATE
			users
		SET
			deleted_at= NULL
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
