package auth

import (
	"context"
	"database/sql"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/entities"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, userId int, refreshToken string) error
	Update(ctx context.Context, tx *sql.Tx, userId int, refreshToken string) error
	Delete(ctx context.Context, tx *sql.Tx, userId int) error
	FindRefreshTokenByUserId(ctx context.Context, tx *sql.Tx, userId int) (entities.Auth, error)
}
