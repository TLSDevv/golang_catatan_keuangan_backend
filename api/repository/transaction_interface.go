package repository

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
)

type TransactionRepository interface {
	ListByUser(ctx context.Context, tx *sql.Tx, limit int, page int, userId int) ([]domain.Transaction, error)
	GetByID(ctx context.Context, tx *sql.Tx, transactionId int) (domain.Transaction, error)
	Store(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) error
	Update(ctx context.Context, tx *sql.Tx, transactionId int, transaction domain.Transaction) error
	Delete(ctx context.Context, tx *sql.Tx, transactionId int) error
}
