package user

type RequestUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
	Role     int    `json:"role" validate:"required"`
}

type RequestRegisterUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
}

type RequestUpdateUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Fullname string `json:"fullname" validate:"required"`
}

type RequestUpdatePassword struct {
	Password string `json:"password" validate:"required"`
}

type RequestUpdateRole struct {
	Role int `json:"role" validate:"required"`
}

//validation Error Response json validate must be string.
type ValidationErrorResponse struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Role     string `json:"role,omitempty"`
}

type ValidationRegisterUserErrorResponse struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Fullname string `json:"fullname,omitempty"`
}
