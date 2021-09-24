package helper

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

func ToUserResponse(user domain.User) *web.UserResponse {
	return &web.UserResponse{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
		Name:     user.Name,
	}
}

func ToUserCreate(user web.UserCreateRequest) domain.User {
	password := HashPassword(user.Password)
	return domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: password,
		Name:     user.Name,
	}
}

func ToUserUpdate(user web.UserUpdateRequest) domain.User {
	password := HashPassword(user.Password)
	return domain.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: password,
		Name:     user.Name,
	}
}
