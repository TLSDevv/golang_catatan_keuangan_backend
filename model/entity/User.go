package entity

import (
	"errors"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
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

	user.Password = pkg.PasswordToHash(user.Password)

	return nil
}

func (u User) CheckPassword(password string) error {
	if pkg.ComparePassword(password, u.Password) {
		return nil
	}

	return errors.New("Wrong username and password!")
}

// func (u User) Update(userDto dto.UserRequest) {
// 	if len(userDto.Username) != 0 {
// 		u.Username = userDto.Username
// 	}

// 	if len(userDto.Email) != 0 {
// 		u.Email = userDto.Email
// 	}

// 	if len(userDto.Password) != 0 {
// 		u.Password = userDto.Password
// 	}

// 	if len(userDto.Fullname) != 0 {
// 		u.Fullname = userDto.Fullname
// 	}

// 	u.UpdatedAt = time.Now()
// }
