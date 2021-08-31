package repository

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	Store(ctx context.Context, tx *sql.Tx, user domain.User) error
	Update(ctx context.Context, tx *sql.Tx, id int, user domain.User) error
}
