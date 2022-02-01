package repository

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/entity"
)

type TransactionRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.Transaction, error)
	FindById(ctx context.Context, tx *sql.Tx, trcId int) (entity.Transaction, error)
	Create(ctx context.Context, tx *sql.Tx, trc entity.Transaction) error
	Update(ctx context.Context, tx *sql.Tx, trc entity.Transaction) error
	Delete(ctx context.Context, tx *sql.Tx, trcId int) error
	Restore(ctx context.Context, tx *sql.Tx, trcId int) error
	Purge(ctx context.Context, tx *sql.Tx, trcId int) error
}
