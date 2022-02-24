package auth

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Service interface {
	Login(ctx context.Context, login entities.Login) (entities.Token, error)
	Logout(ctx context.Context) error
	Refresh(ctx context.Context, userId int) (entities.Token, error)
}
