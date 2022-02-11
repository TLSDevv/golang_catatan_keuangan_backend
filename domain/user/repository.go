package user

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Repository interface {
	Create(ctx context.Context, user entities.User) error
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, user entities.User) error
	Purge(ctx context.Context, user entities.User) error
	Restore(ctx context.Context, user entities.User) error
	FindById(ctx context.Context, userId int) (entities.User, error)
	FindByUsername(ctx context.Context, username string) (entities.User, error)
	List(ctx context.Context) ([]entities.User, error)
}
