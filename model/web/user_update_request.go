package web

type UserUpdateRequest struct {
	Id       uint8  `validate:"required" json:"id"`
	Username string `validate:"required,max=100,min=1" json:"username"`
	Email    string `validate:"required,max=100,min=1" json:"email"`
	Password string `validate:"required,max=100,min=1" json:"password"`
	Name     string `validate:"required,max=100,min=1" json:"name"`
}
