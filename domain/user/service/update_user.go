package service

import (
	"context"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (service UserService) UpdateUser(ctx context.Context, id int, userRequest entities.User) error {
	user, err := service.UserRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	user.UpdateUser(userRequest)

	err = service.UserRepository.UpdateUser(ctx, id, user)

	if err != nil {
		return err
	}

	return nil
}
