package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Repository interface {
	GetTransactions(ctx context.Context) ([]entities.Transaction, error)         // only admin
	GetActiveTransactions(ctx context.Context) ([]entities.Transaction, error)   // only admin
	GetInActiveTransactions(ctx context.Context) ([]entities.Transaction, error) // only admin

	GetByID(ctx context.Context, transactionID int) (entities.Transaction, error)       // only admin
	GetActiveByID(ctx context.Context, transactionID int) (entities.Transaction, error) // admin, user

	GetTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)         // admin,user
	GetActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)   // admin,user
	GetInActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) // admin
	GetTodayTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)    // admin, user
	// GetTodayActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)       // admin, user
	// GetTodayInActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)     // admin
	// GetThisWeekActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)    // admin, user
	// GetThisMonthInActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error) // admin, user
	// GetThisYearActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)    // admin, user
	// GetThisYearInActiveTransactionsByUserID(ctx context.Context, userID int) ([]entities.Transaction, error)  // admin, user

	Create(ctx context.Context, trc entities.Transaction) error          // admin,user
	Update(ctx context.Context, trc entities.Transaction, tID int) error // admin,user
	Delete(ctx context.Context, transactionID int) error                 // admin,user
	Restore(ctx context.Context, transactionID int) error                // only admin
	Purge(ctx context.Context, transactionID int) error                  // only admin

	// CheckUser(ctx context.Context, userID int) (bool, error)
	CheckTransactionByID(ctx context.Context, tID int) (bool, error)
}
