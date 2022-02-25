package service

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/domain/user"
)

type UserService struct {
	UserRepository user.Repository
}

func NewUserService(userRepo user.Repository) UserService {
	return UserService{
		UserRepository: userRepo,
	}
}
