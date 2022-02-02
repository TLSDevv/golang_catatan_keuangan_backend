package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Service interface {
	Create(ctx context.Context, user entities.UserInput) error
	Update(ctx context.Context, id int, user entities.UserInput) error
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (user entities.UserResponse, err error)
	List(ctx context.Context) ([]entities.UserResponse, error)
}
