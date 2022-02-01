package user

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Repository interface {
	Create(ctx context.Context, tx *sql.Tx, user entities.User) error
	Update(ctx context.Context, tx *sql.Tx, user entities.User) error
	Delete(ctx context.Context, tx *sql.Tx, user entities.User) error
	Purge(ctx context.Context, tx *sql.Tx, user entities.User) error
	Restore(ctx context.Context, tx *sql.Tx, user entities.User) error
	FindById(ctx context.Context, tx *sql.Tx, userId int) (entities.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (entities.User, error)
	List(ctx context.Context, tx *sql.Tx) ([]entities.User, error)
}
