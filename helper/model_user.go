package helper

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Gender:   user.Gender,
		Age:      user.Age,
		Job:      user.Job,
	}
}

func ToUserCreate(user web.UserCreateRequest) domain.User {
	return domain.User{
		Username: user.Username,
		Name:     user.Name,
		Gender:   user.Gender,
		Age:      user.Age,
		Job:      user.Job,
	}
}

func ToUserUpdate(user web.UserUpdateRequest) domain.User {
	return domain.User{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Gender:   user.Gender,
		Age:      user.Age,
		Job:      user.Job,
	}
}
