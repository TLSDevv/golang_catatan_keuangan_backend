package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (t Transaction) GetTransactions(ctx context.Context) ([]entities.Transaction, error) {
	transactions, err := t.tr.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
