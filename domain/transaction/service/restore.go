package service

import "context"

func (t Transaction) Restore(ctx context.Context, tID int) error {
	err := t.tr.Restore(ctx, tID)
	if err != nil {
		return err
	}

	return nil
}
