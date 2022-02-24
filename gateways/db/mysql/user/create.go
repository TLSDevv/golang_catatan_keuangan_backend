package user

import (
	"context"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
)

func (u UserRepository) Create(ctx context.Context, user entities.User) error {
	errors := []string{}

	usernameIsExist, err := u.usernameIsExist(ctx, user.Username)
	emailIsExist, err := u.emailIsExist(ctx, user.Email)

	if err != nil {
		return err
	}

	if usernameIsExist {
		errors = append(errors, "username already exist")
	}

	if emailIsExist {
		errors = append(errors, "email already exist")
	}

	if len(errors) != 0 {
		return helper.NewErrors("validation Error", errors)
	}

	sql := `
		INSERT INTO
			users(
				username,
				email,
				password,
				fullname,
				role,
				updated_at)
		VALUES( ?, ?, ?, ?, ?, ?)`

	_, err = u.DB.ExecContext(ctx, sql,
		user.Username,
		user.Email,
		user.Password,
		user.Fullname,
		user.Role,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) usernameIsExist(ctx context.Context, username string) (exist bool, err error) {
	sql := `
		SELECT
			username
		FROM
			users
		WHERE username = ?`

	rows, err := u.DB.QueryContext(ctx, sql,
		username,
	)

	if err != nil {
		return false, err
	}

	if rows.Next() {
		exist = true
	}

	return
}

func (u UserRepository) emailIsExist(ctx context.Context, email string) (exist bool, err error) {
	sql := `
		SELECT
			email
		FROM
			users
		WHERE email = ?`

	rows, err := u.DB.QueryContext(ctx, sql,
		email,
	)

	if err != nil {
		return false, err
	}

	if rows.Next() {
		exist = true
	}

	return
}
