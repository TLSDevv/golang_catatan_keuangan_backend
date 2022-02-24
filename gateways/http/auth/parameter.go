package auth

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//validation Error Response json validate must be string.
type ValidationLoginResponse struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
