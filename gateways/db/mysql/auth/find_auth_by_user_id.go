package auth

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

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
