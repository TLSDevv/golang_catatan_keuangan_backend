package web

type UserCreateRequest struct {
	Username string `validate:"required,max=100,min=1" json:"username"`
	Name     string `validate:"required,max=100,min=1" json:"name"`
	Gender   string `validate:"required,min=2" json:"gender"`
	Age      int    `json:"age"`
	Job      string `validate:"required,max=100,min=1" json:"job"`
}
