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

// request input
type UserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func (user UserInput) Validate() error {
	if len(user.Username) == 0 {
		return ErrUnameRequired
	}

	if len(user.Email) == 0 {
		return ErrEmailRequired
	}

	if len(user.Password) == 0 {
		return ErrPassRequired
	}

	if len(user.Fullname) == 0 {
		return ErrFNameRequired
	}

	return nil
}

func (user User) Validate() error {
	if len(user.Username) == 0 {
		return ErrUnameRequired
	}

	if len(user.Email) == 0 {
		return ErrEmailRequired
	}

	if len(user.Password) == 0 {
		return ErrPassRequired
	}

	if len(user.Fullname) == 0 {
		return ErrFNameRequired
	}

	return nil
}

func (u User) CheckPassword(password string) error {
	if pkg.ComparePassword(password, u.Password) {
		return nil
	}

	return errors.New("Wrong username and password!")
}

func (u User) Update(ui UserInput) {
	if len(ui.Username) != 0 {
		u.Username = ui.Username
	}

	if len(ui.Email) != 0 {
		u.Email = ui.Email
	}

	if len(ui.Password) != 0 {
		u.Password = pkg.PasswordToHash(ui.Password)
	}

	if len(ui.Fullname) != 0 {
		u.Fullname = ui.Fullname
	}

	u.UpdatedAt = time.Now()
}
