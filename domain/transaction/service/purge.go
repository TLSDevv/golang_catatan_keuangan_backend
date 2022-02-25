package service

import "context"

func (t Transaction) Purge(ctx context.Context, tID int) error {
	err := t.tr.Purge(ctx, tID)
	if err != nil {
		return err
	}

	return nil
}
