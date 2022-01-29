package entity

import (
	"errors"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/dto"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Fullname  string    `json:"fullname"`
	createdAt time.Time `json:"created_at"`
	updatedAt time.Time `json:"updated_at"`
	deletedAt time.Time `json:"deleted_at"`
}

func (user User) Validate() error {
	if len(user.Username) == 0 {
		return errors.New("Username required")
	}

	if len(user.Email) == 0 {
		return errors.New("Email required")
	}

	if len(user.Password) == 0 {
		return errors.New("Password required")
	}

	if len(user.Fullname) == 0 {
		return errors.New("Email required")
	}

	// password needs to hash before save

	return nil
}

func (u User) ToDTO() dto.User {
	return dto.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.Fullname,
	}
}

func (u User) CheckPassword(password string) error {
	if u.Password != password {
		return errors.New("Wrong username and password!")
	}
	return nil
}
