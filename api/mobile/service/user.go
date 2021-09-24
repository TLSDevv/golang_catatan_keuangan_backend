package service

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/exception"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/helper"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
	"github.com/go-playground/validator"
)

type UserService struct {
	DB *sql.DB
	//todo fix validate
	Validate *validator.Validate
}

var (
	//Register Repo
	userRepo = repository.NewUserRepository()
)

func NewUserService(db *sql.DB, v *validator.Validate) UserServiceInterface {
	return &UserService{
		DB:       db,
		Validate: v,
	}
}

type UserServiceInterface interface {
	GetUser(ctx context.Context, userId int) *web.UserResponse
	CreateUser(ctx context.Context, userRequest web.UserCreateRequest)
	UpdateUser(ctx context.Context, userRequest web.UserUpdateRequest)
}

func (u *UserService) GetUser(ctx context.Context, userId int) *web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicIfError(err)

	defer func() {
		helper.CommitOrRollback(tx)
	}()

	user, err := userRepo.GetByID(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

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

	err = userRepository.Store(ctx, tx, user)
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

	err = userRepository.Update(ctx, tx, user)
	helper.PanicIfError(err)
}
