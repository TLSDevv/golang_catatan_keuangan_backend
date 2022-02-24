package service

import (
	"context"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

func (service UserService) FindById(ctx context.Context, id int) (entities.UserResponse, error) {
	user, err := service.UserRepository.FindById(ctx, id)

	if err != nil {
		return entities.UserResponse{}, err
	}

	if user.ID == 0 {
		return entities.UserResponse{}, errors.New("User Not Found")
	}

	return entities.UserToUserResponse(user), nil
}
