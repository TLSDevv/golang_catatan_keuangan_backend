package service

import (
	"context"
	"errors"
)

func (service UserService) UpdateRole(ctx context.Context, id int, role int) error {
	user, err := service.UserRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	err = service.UserRepository.UpdateRole(ctx, id, role)

	if err != nil {
		return err
	}

	return nil
}
