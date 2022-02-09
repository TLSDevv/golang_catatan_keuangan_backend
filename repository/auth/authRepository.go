package auth

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type AuthRepository struct {
}

func NewAuthRepository() AuthRepository {
	return AuthRepository{}
}

func (repository AuthRepository) FindRefreshTokenByUserId(ctx context.Context, tx *sql.Tx, userId int) (entities.Auth, error) {
	sql := `
		SELECT
			refresh_token
		FROM
			auths
		WHERE
			user_id=$1`

	rows, err := tx.QueryContext(ctx, sql,
		userId,
	)

	if err != nil {
		return entities.Auth{}, err
	}

	auth := entities.Auth{}

	if rows.Next() {
		err := rows.Scan(
			&auth.RefreshToken,
		)

		if err != nil {
			return entities.Auth{}, err
		}
	}

	return auth, nil
}

func (repository AuthRepository) Save(ctx context.Context, tx *sql.Tx, userId int, refreshToken string) error {
	sql := `
		INSERT INTO
			auths(
				user_id,
				refresh_token)
			VALUES($1, $2)`

	_, err := tx.ExecContext(ctx, sql,
		userId,
		refreshToken,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository AuthRepository) Update(ctx context.Context, tx *sql.Tx, userId int, refreshToken string) error {
	sql := `
		UPDATE
			auths
		SET
			refresh_token=$2,
		WHERE
			user_id=$1`

	_, err := tx.ExecContext(ctx, sql,
		userId,
		refreshToken,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repository AuthRepository) Delete(ctx context.Context, tx *sql.Tx, userId int) error {
	sql := `
		DELETE FROM
			auths
		WHERE
			user_id=$1`

	_, err := tx.ExecContext(ctx, sql,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}
