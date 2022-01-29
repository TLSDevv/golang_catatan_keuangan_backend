package dto

import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func (u UserRequest) Validate() error {
	if len(u.Username) == 0 {
		return errors.New("Username required")
	}

	if len(u.Email) == 0 {
		return errors.New("Email required")
	}

	if len(u.Password) == 0 {
		return errors.New("Password required")
	}

	if len(u.Fullname) == 0 {
		return errors.New("Email required")
	}
	return nil
}
