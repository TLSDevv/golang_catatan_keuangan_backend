package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
)

type TransactionRepository interface {
	List(ctx context.Context, tx *sql.Tx, limit int, page int) ([]domain.Transaction, error)
	GetByID(ctx context.Context, tx *sql.Tx, id int) (domain.Transaction, error)
	Store(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) error
	Update(ctx context.Context, tx *sql.Tx, id int, transaction domain.Transaction) error
	Delete(ctx context.Context, tx *sql.Tx, id int, deleteAt time.Time) error
}
