package service

import "context"

func (t Transaction) CheckTransactionByID(ctx context.Context, tID int) (bool, error) {
	tExist, err := t.tr.CheckTransactionByID(ctx, tID)
	if err != nil {
		return false, err
	}

	return tExist, nil
}
