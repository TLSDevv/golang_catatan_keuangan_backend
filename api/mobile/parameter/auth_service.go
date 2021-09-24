package parameter

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	TokenData TokenData         `json:"token"`
	UserData  *web.UserResponse `json:"user"`
}

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	// Is2FAEnabled bool   `json:"two_fa_enabled"`
	// IsActive     bool   `json:"is_active"`
	// ExpiredAt    string `json:"expired_at"`
}

type LogoutRequest struct {
}
type LogoutResponse struct {
}