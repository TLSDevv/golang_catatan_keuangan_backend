package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
)

type userRepo struct {
}

var structureUserStore string = `username,name,gender,age,job,created_at,updated_at`
var structureUserUpdate string = `username,name,gender,age,job,updated_at`
var structureUser string = `id,username,name,gender,age,job`

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (u *userRepo) Store(ctx context.Context, tx *sql.Tx, user domain.User) error {
	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = user.CreatedAt
	sql := `INSERT INTO USERS (
		` + structureUserStore + `)
	values ($1,$2,$3,$4,$5,$6,$7)`

	_, err := tx.ExecContext(
		ctx, sql,
		user.Username,
		user.Name,
		user.Gender,
		user.Age,
		user.Job,
		user.CreatedAt,
		user.UpdatedAt,
	)
	helper.PanicIfError(err)

	return nil
}

func (u *userRepo) Update(ctx context.Context, tx *sql.Tx, user domain.User) error {
	user.UpdatedAt = time.Now().Local()
	sql := `UPDATE INTO USERS (
		` + structureUserUpdate + `)
	values ($1,$2,$3,$4,$5,$6) WHERE id=$7`

	_, err := tx.ExecContext(
		ctx, sql,
		user.Username,
		user.Name,
		user.Gender,
		user.Age,
		user.Job,
		user.UpdatedAt,
		user.Id,
	)
	helper.PanicIfError(err)

	return nil
}

func (t *userRepo) GetByID(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	sql := `SELECT ` + structureUser + ` FROM users WHERE id=$1 AND deleted_at IS NULL`
	rows, err := tx.QueryContext(ctx, sql, id)

	helper.PanicIfError(err)

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Name,
			&user.Gender,
			&user.Age,
			&user.Job,
		)
		helper.PanicIfError(err)

		return user, nil
	} else {
		return user, errors.New("User Not Found")
	}
}
