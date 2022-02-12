package auth

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Repository interface {
	Save(ctx context.Context, userId int, refreshToken string) error
	Update(ctx context.Context, userId int, refreshToken string) error
	Delete(ctx context.Context, userId int) error
	FindRefreshTokenByUserId(ctx context.Context, userId int) (entities.Auth, error)
}
