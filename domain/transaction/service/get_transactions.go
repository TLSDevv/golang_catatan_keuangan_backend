package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (t Transaction) GetTransactions(ctx context.Context) ([]entities.Transaction, error) {
	uID := ctx.Value("user_id").(int)
	transactions, err := t.tr.GetActiveTransactionsByUserID(ctx, uID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
