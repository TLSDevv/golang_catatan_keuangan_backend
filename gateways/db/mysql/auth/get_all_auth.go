package auth

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (repository AuthRepository) GetAllAuth(ctx context.Context) ([]entities.Auth, error) {
	sql := `
		SELECT
			user_id, refresh_token
		FROM
			auths`

	rows, err := repository.DB.QueryContext(ctx, sql)

	if err != nil {
		return nil, err
	}

	auths := []entities.Auth{}

	if rows.Next() {
		auth := entities.Auth{}
		err := rows.Scan(
			&auth.UserId,
			&auth.RefreshToken,
		)

		if err != nil {
			return nil, err
		}

		auths = append(auths, auth)
	}

	return auths, nil
}
