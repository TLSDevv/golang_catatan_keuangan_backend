package service

import (
	"context"
)

func (t Transaction) CheckUser(ctx context.Context, userID int) (bool, error) {
	userExist, err := t.repo.CheckUser(ctx, userID)
	if err != nil {
		return false, err
	}

	return userExist, nil
}
