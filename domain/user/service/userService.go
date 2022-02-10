package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

type UserService struct {
	UserRepository user.Repository
	DB             *sql.DB
}

func NewUserService(userRepo user.Repository, db *sql.DB) UserService {
	return UserService{
		UserRepository: userRepo,
		DB:             db,
	}
}

func (service UserService) Create(ctx context.Context, userRequest entities.UserInput) error {
	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	password := pkg.PasswordToHash(userRequest.Password)

	user := entities.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: password,
		Fullname: userRequest.Fullname,
	}

	err = service.UserRepository.Create(ctx, tx, user)

	if err != nil {
		return err
	}

	return nil
}

func (service UserService) Update(ctx context.Context, id int, userRequest entities.UserInput) error {
	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	user, err := service.UserRepository.FindById(ctx, tx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	user.Update(userRequest)

	err = service.UserRepository.Update(ctx, tx, user)

	if err != nil {
		return err
	}

	return nil
}

func (service UserService) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.Begin()

	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	user, err := service.UserRepository.FindById(ctx, tx, id)

	if err != nil {
		return err
	}

	if user.ID == 0 {
		return errors.New("User Not Found")
	}

	err = service.UserRepository.Delete(ctx, tx, user)

	if err != nil {
		return err
	}

	return nil
}

func (service UserService) FindById(ctx context.Context, id int) (entities.UserResponse, error) {
	tx, err := service.DB.Begin()

	if err != nil {
		return entities.UserResponse{}, nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	user, err := service.UserRepository.FindById(ctx, tx, id)

	if err != nil {
		return entities.UserResponse{}, err
	}

	if user.ID == 0 {
		return entities.UserResponse{}, errors.New("User Not Found")
	}

	return entities.UserToUserResponse(user), nil
}

func (service UserService) List(ctx context.Context) ([]entities.UserResponse, error) {
	tx, err := service.DB.Begin()

	if err != nil {
		return []entities.UserResponse{}, nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	users, err := service.UserRepository.List(ctx, tx)

	return entities.UsersToUsersResponse(users), nil
}
