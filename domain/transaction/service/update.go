package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (t Transaction) Update(ctx context.Context, input entities.TransactionInput, tID int) error {
	tr, err := t.tr.GetByID(ctx, tID)
	if err != nil {
		return err
	}

	tr.Update(input)

	err = t.tr.Update(ctx, tr, tID)
	if err != nil {
		return err
	}

	return nil
}
