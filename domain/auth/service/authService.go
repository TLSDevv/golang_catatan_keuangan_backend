package service

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/auth"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
)

type AuthService struct {
	UserRepo user.Repository
	AuthRepo auth.Repository
}

func NewAuthService(authRepo auth.Repository, userRepo user.Repository) AuthService {
	return AuthService{
		AuthRepo: authRepo,
		UserRepo: userRepo,
	}
}
