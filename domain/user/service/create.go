package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

func (service UserService) Create(ctx context.Context, userRequest entities.User) error {
	password := pkg.PasswordToHash(userRequest.Password)

	userRequest.Password = password

	err := service.UserRepository.Create(ctx, userRequest)

	if err != nil {
		return err
	}

	return nil
}
