package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

func (t Transaction) Create(ctx context.Context, input entities.TransactionInput) error {
	transactionAt, err := pkg.StringDateToDateTime(input.TransactionAt)
	if err != nil {
		return err
	}

	transaction, err := entities.NewTransaction(
		ctx.Value("user_id").(int),
		input.TransactionName,
		input.Category,
		input.TransactionType,
		input.Amount,
		*transactionAt,
	)

	err = t.tr.Create(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}
