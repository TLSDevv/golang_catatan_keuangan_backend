package service

import (
	"context"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

type UserService struct {
	UserRepository user.Repository
}

func NewUserService(userRepo user.Repository) UserService {
	return UserService{
		UserRepository: userRepo,
	}
}

func (service UserService) Create(ctx context.Context, userRequest entities.UserInput) error {
	password := pkg.PasswordToHash(userRequest.Password)

	user := entities.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: password,
		Fullname: userRequest.Fullname,
	}

	err := service.UserRepository.Create(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (service UserService) Update(ctx context.Context, id int, userRequest entities.UserInput) error {
	user, err := service.UserRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	user.Update(userRequest)

	err = service.UserRepository.Update(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (service UserService) Delete(ctx context.Context, id int) error {
	user, err := service.UserRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	err = service.UserRepository.Delete(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

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

func (service UserService) List(ctx context.Context) ([]entities.UserResponse, error) {
	users, err := service.UserRepository.List(ctx)

	if err != nil {
		return []entities.UserResponse{}, nil
	}

	return entities.UsersToUsersResponse(users), nil
}
