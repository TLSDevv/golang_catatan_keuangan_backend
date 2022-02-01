package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Service interface {
	Create(ctx context.Context, user entities.UserResponse) error
	Update(ctx context.Context, id int, user entities.UserResponse) error
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (entities.User, error)
	List(ctx context.Context) ([]entities.User, error)
}
