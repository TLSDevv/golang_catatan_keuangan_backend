package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Service interface {
	GetTransactions(ctx context.Context) ([]entities.Transaction, error)
	GetByID(ctx context.Context, transactionID int) (*entities.Transaction, error)
	// GetTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)
	// GetTodayTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)
	Create(ctx context.Context, trc entities.TransactionInput) error
	CheckUser(ctx context.Context, userID int) (bool, error)
	CheckTransactionByID(ctx context.Context, tID int) (bool, error)
	Update(ctx context.Context, trc entities.TransactionInput) error
	// Delete(ctx context.Context, tID int) error
	// Restore(ctx context.Context, tID int) error
	// Purge(ctx context.Context, tID int) error
}
