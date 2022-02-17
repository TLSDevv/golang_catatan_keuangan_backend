package entities

import (
	"errors"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

var (
	ErrUnameRequired     = errors.New("Username is required")
	ErrUnameAlreadyExist = errors.New("Username is already in use")
	ErrEmailRequired     = errors.New("Email is required")
	ErrEmailAlreadyExist = errors.New("Email is already in use")
	ErrPassRequired      = errors.New("Password is required")
	ErrPassInvalid       = errors.New("Password is required")
	ErrFNameRequired     = errors.New("Fullname is required")

	UserDomainErrors = []error{
		ErrUnameRequired,
		ErrUnameAlreadyExist,
		ErrEmailRequired,
		ErrEmailAlreadyExist,
		ErrPassRequired,
		ErrPassInvalid,
		ErrFNameRequired,
	}
)

// entity
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

func (u *User) UpdateUser(userRequest User) {
	u.Username = userRequest.Username
	u.Email = userRequest.Email
	u.Fullname = userRequest.Fullname
}

func (u *User) UpdatePassword(password string) {
	u.Password = password
}

func UserToUserResponse(u User) UserResponse {
	return UserResponse{
		u.ID,
		u.Username,
		u.Email,
		u.Fullname,
	}
}

func UsersToUsersResponse(u []User) []UserResponse {
	users := []UserResponse{}

	for _, value := range u {
		user := UserResponse{
			value.ID,
			value.Username,
			value.Email,
			value.Fullname,
		}

		users = append(users, user)
	}

	return users
}

// response data
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

func (u User) CheckPassword(password string) error {
	if pkg.ComparePassword(password, u.Password) {
		return nil
	}

	return errors.New("Wrong username and password!")
}
