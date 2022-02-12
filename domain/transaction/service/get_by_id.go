package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (t Transaction) GetByID(ctx context.Context, transactionID int) (*entities.Transaction, error) {
	transaction, err := t.repo.GetByID(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
