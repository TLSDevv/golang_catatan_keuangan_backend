package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
)

type CategoryRepository interface {
	ListByUser(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Category, error)
	GetByID(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	Store(ctx context.Context, tx *sql.Tx, category domain.Category) error
	Update(ctx context.Context, tx *sql.Tx, id int, transaction domain.Category) error
	Delete(ctx context.Context, tx *sql.Tx, id int, deleteAt time.Time) error
}
