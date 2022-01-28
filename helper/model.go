package helper

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/dto"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/entity"
)

func UserToDTO(u entity.User) dto.User {
	return dto.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.Fullname,
	}
}

func UsersToDTO(u []entity.User) []dto.User {
	users := []dto.User{}

	for _, value := range u {
		user := UserToDTO(value)
		users = append(users, user)
	}

	return users
}

func UserDTOToUser(u dto.User) entity.User {
	return entity.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.Fullname,
	}
}
