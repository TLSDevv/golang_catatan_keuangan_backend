package transaction

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Service interface {
	FindAll(ctx context.Context) ([]entities.TransactionResponse, error)
}
