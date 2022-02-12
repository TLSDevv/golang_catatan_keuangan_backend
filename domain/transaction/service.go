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
	Create(ctx context.Context, trc entities.CreateTransactionInput) error
	CheckUser(ctx context.Context, userID int) (bool, error)
	// Update(ctx context.Context, trc entities.Transaction) error
	// Delete(ctx context.Context, transactionID int) error
	// Restore(ctx context.Context, transactionID int) error
	// Purge(ctx context.Context, transactionID int) error
}
