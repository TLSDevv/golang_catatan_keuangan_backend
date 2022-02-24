package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (service UserService) List(ctx context.Context) ([]entities.UserResponse, error) {
	users, err := service.UserRepository.List(ctx)

	if err != nil {
		return []entities.UserResponse{}, nil
	}

	return entities.UsersToUsersResponse(users), nil
}
