package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (u UserRepository) List(ctx context.Context) ([]entities.User, error) {
	sql := `
		SELECT
			id, username, email, password, fullname, role, created_at, updated_at,
		FROM
			users
		WHERE
		deleted_at = null`

	rows, err := u.DB.QueryContext(ctx, sql)

	if err != nil {
		return []entities.User{}, err
	}

	defer rows.Close()

	users := []entities.User{}

	if rows.Next() {
		user := entities.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Fullname,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return []entities.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
