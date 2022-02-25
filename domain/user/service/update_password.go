package service

import (
	"context"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

func (service UserService) UpdatePassword(ctx context.Context, id int, password string) error {
	user, err := service.UserRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	hashPassword := pkg.PasswordToHash(password)

	err = service.UserRepository.UpdatePassword(ctx, id, hashPassword)

	if err != nil {
		return err
	}

	return nil
}
