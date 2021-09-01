package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type UserService struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepo repository.UserRepository, db *sql.DB) UserServiceInterface {
	return &UserService{
		UserRepository: userRepo,
		DB:             db,
	}
}

func (u *UserService) GetUser(ctx context.Context, userId int) web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	user, err := u.UserRepository.GetByID(ctx, tx, userId)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}
func (u *UserService) CreateUser(ctx context.Context, userRequest web.UserCreateRequest) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	user := helper.ToUser(userRequest)

	err = u.UserRepository.Store(ctx, tx, user)
	helper.PanicIfError(err)
}
func (u *UserService) UpdateUser(ctx context.Context, userId int, userRequest web.UserCreateRequest) {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	user := helper.ToUser(userRequest)

	err = u.UserRepository.Update(ctx, tx, userId, user)
	helper.PanicIfError(err)
}
