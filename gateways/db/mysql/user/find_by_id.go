package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (u UserRepository) FindById(ctx context.Context, userId int) (entities.User, error) {
	sql := `
		SELECT
			id, username, email, password, fullname, role, created_at, updated_at, deleted_at
		FROM
			users
		WHERE
			id=?
			AND deleted_at IS NULL`

	rows, err := u.DB.QueryContext(ctx, sql,
		userId,
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
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return entities.User{}, err
		}
	}

	return user, nil
}
