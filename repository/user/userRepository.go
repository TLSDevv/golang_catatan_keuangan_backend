package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

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

func (u UserRepository) Update(ctx context.Context, user entities.User) error {
	sql := `
		UPDATE
			users
		SET
			username=?,
			email=?,
			password=?,
			fullname=?,
			updated_at=?
		WHERE
			id=?`

	_, err := u.DB.ExecContext(ctx, sql,
		user.Username,
		user.Email,
		user.Password,
		user.Fullname,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Purge(ctx context.Context, user entities.User) error {
	sql := `
		DELETE FROM
			users
		WHERE
			id=?`

	_, err := u.DB.ExecContext(ctx, sql,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Delete(ctx context.Context, user entities.User) error {
	sql := `
		UPDATE
			users
		SET
			deleted_at=?
		WHERE
			id=?`

	deletedAt := time.Now()

	_, err := u.DB.ExecContext(ctx, sql,
		deletedAt,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Restore(ctx context.Context, user entities.User) error {
	sql := `
		UPDATE
			users
		SET
			deleted_at= NULL
		WHERE
			id=?`

	_, err := u.DB.ExecContext(ctx, sql,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) FindById(ctx context.Context, userId int) (entities.User, error) {
	sql := `
		SELECT
			id, username, email, password, fullname, created_at, updated_at, deleted_at
		FROM
			users
		WHERE
			id=?`

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

func (u UserRepository) FindByUsername(ctx context.Context, username string) (entities.User, error) {
	sql := `
		SELECT
			id, username, email, password, fullname, created_at, updated_at
		FROM
			users
		WHERE
			username=?`

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

func (u UserRepository) List(ctx context.Context) ([]entities.User, error) {
	sql := `
		SELECT
			id, username, email, password, fullname, created_at, updated_at, deleted_at
		FROM
			users`

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
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return []entities.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u UserRepository) CheckUser(ctx context.Context, userID int) (bool, error) {
	var result entities.CountResult

	query := `SELECT COUNT(*) as total FROM users WHERE id = ?`
	err := u.DB.QueryRowContext(ctx, query, userID).Scan(&result.Total)
	if err != nil {
		return false, err
	}

	return result.Total > 0, nil
}
