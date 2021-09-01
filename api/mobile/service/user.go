package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
	"github.com/go-playground/validator"
)

type UserService struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepo repository.UserRepository, db *sql.DB, v *validator.Validate) UserServiceInterface {
	return &UserService{
		UserRepository: userRepo,
		DB:             db,
		Validate:       v,
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
	err := u.Validate.Struct(userRequest)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	user := helper.ToUserCreate(userRequest)

	err = u.UserRepository.Store(ctx, tx, user)
	helper.PanicIfError(err)
}
func (u *UserService) UpdateUser(ctx context.Context, userRequest web.UserUpdateRequest) {
	err := u.Validate.Struct(userRequest)
	helper.PanicIfError(err)

	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	user := helper.ToUserUpdate(userRequest)

	err = u.UserRepository.Update(ctx, tx, user)
	helper.PanicIfError(err)
}
