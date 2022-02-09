package entities

import "errors"

type UserDetails struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l Login) Validate() error {
	if l.Username == "" {
		return errors.New("Required Username")
	}
	if l.Password == "" {
		return errors.New("Required Password")
	}
	return nil
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Auth struct {
	UserId       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}
