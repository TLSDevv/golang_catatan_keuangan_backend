package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (u UserRepository) FindByUsername(ctx context.Context, username string) (entities.User, error) {
	sql := `
		SELECT
			id, username, email, password, fullname, created_at, updated_at
		FROM
			users
		WHERE
			username=? AND deleted_at=null`

	rows, err := u.DB.QueryContext(ctx, sql,
		username,
	)

	if err != nil {
		return entities.User{}, err
	}

	defer rows.Close()

	user := entities.User{}

	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Fullname,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return entities.User{}, err
		}
	}

	return user, nil
}
