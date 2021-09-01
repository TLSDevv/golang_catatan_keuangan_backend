package web

type UserCreateRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Job      string `json:"job"`
}
