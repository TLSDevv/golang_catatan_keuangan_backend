package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Service interface {
	GetTransactions(ctx context.Context) ([]entities.Transaction, error)
	GetByID(ctx context.Context, transactionID int) (*entities.Transaction, error)
}
