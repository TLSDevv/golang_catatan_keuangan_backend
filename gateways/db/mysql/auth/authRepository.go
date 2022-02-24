package auth

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return AuthRepository{
		DB: db,
	}
}

func (repository AuthRepository) FindAuthByUserId(ctx context.Context, userId int) (entities.Auth, error) {
	sql := `
		SELECT
			user_id, refresh_token
		FROM
			auths
		WHERE
			user_id=?`

	rows, err := repository.DB.QueryContext(ctx, sql,
		userId,
	)

	if err != nil {
		return entities.Auth{}, err
	}

	auth := entities.Auth{}

	if rows.Next() {
		err := rows.Scan(
			&auth.UserId,
			&auth.RefreshToken,
		)

		if err != nil {
			return entities.Auth{}, err
		}
	}

	return auth, nil
}

func (repository AuthRepository) Save(ctx context.Context, userId int, refreshToken string) error {

	auth, err := repository.FindAuthByUserId(ctx, userId)

	if auth.UserId != 0 {
		repository.Update(ctx, userId, refreshToken)

		return nil
	}

	sql := `
		INSERT INTO
			auths(
				user_id,
				refresh_token)
			VALUES(?, ?)`

	_, err = repository.DB.ExecContext(ctx, sql,
		userId,
		refreshToken,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository AuthRepository) Update(ctx context.Context, userId int, refreshToken string) error {
	sql := `
		UPDATE
			auths
		SET
			refresh_token=?,
		WHERE
			user_id=?`

	_, err := repository.DB.ExecContext(ctx, sql,
		refreshToken,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository AuthRepository) Delete(ctx context.Context, userId int) error {
	sql := `
		DELETE FROM
			auths
		WHERE
			user_id=?`

	_, err := repository.DB.ExecContext(ctx, sql,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}
