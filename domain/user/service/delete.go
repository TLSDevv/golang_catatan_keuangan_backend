package service

import (
	"context"
	"errors"
)

func (service UserService) Delete(ctx context.Context, id int) error {
	user, err := service.UserRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	err = service.UserRepository.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
