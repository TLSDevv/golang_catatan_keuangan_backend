package service

import (
	"context"
)

func (t Transaction) Delete(ctx context.Context, tID int) error {
	err := t.tr.Delete(ctx, tID)
	if err != nil {
		return err
	}

	return nil
}
