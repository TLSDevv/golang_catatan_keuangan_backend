package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Service interface {
	Create(ctx context.Context, user entities.User) error
	UpdateUser(ctx context.Context, id int, user entities.User) error
	UpdatePassword(ctx context.Context, id int, password string) error
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (user entities.UserResponse, err error)
	List(ctx context.Context) ([]entities.UserResponse, error)
}
