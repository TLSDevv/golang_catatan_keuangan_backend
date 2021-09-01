package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type UserServiceInterface interface {
	GetUser(ctx context.Context, userId int) web.UserResponse
	CreateUser(ctx context.Context, userRequest web.UserCreateRequest)
	UpdateUser(ctx context.Context, userRequest web.UserUpdateRequest)
}
