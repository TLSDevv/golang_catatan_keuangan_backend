package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (u UserRepository) UpdateUser(ctx context.Context, id int, user entities.User) error {
	sql := `
		UPDATE
			users
		SET
			username=?,
			email=?,
			fullname=?,
			updated_at=?
		WHERE
			id=?`

	_, err := u.DB.ExecContext(ctx, sql,
		user.Username,
		user.Email,
		user.Fullname,
		user.UpdatedAt,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
