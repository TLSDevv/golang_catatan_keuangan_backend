package user

import (
	"context"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (u UserRepository) Create(ctx context.Context, user entities.User) error {
	sql := `
		INSERT INTO
			users(
				username,
				email,
				password,
				fullname,
				updated_at)
		VALUES( ?, ?, ?, ?, ?)`

	_, err := u.DB.ExecContext(ctx, sql,
		user.Username,
		user.Email,
		user.Password,
		user.Fullname,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
