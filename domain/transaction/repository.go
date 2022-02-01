package transaction

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

// repository contract
type Repository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]entities.Transaction, error)
	FindById(ctx context.Context, tx *sql.Tx, trcId int) (entities.Transaction, error)
	Create(ctx context.Context, tx *sql.Tx, trc entities.Transaction) error
	Update(ctx context.Context, tx *sql.Tx, trc entities.Transaction) error
	Delete(ctx context.Context, tx *sql.Tx, trcId int) error
	Restore(ctx context.Context, tx *sql.Tx, trcId int) error
	Purge(ctx context.Context, tx *sql.Tx, trcId int) error
}
