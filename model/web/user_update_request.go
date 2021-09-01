package web

type UserUpdateRequest struct {
	Id       uint8  `validate:"required" json:"id"`
	Username string `validate:"required,max=100,min=1" json:"username"`
	Name     string `validate:"required,max=100,min=1" json:"name"`
	Gender   string `validate:"required,min=2" json:"gender"`
	Age      int    `validate:"required,min=1" json:"age"`
	Job      string `validate:"required,max=100,min=1" json:"job"`
}
