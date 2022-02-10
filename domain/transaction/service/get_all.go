package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (s Transaction) GetTransactions(ctx context.Context) ([]entities.Transaction, error) {
	transactions, err := s.repo.GetTransactions(ctx)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
