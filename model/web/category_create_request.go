package web

type CategoryCreateRequest struct {
	UserId       uint8  `validate:"required" json:"user_id"`
	NameCategory string `validate:"required,max=50,min=1" json:"name_category"`
	Description  string `validate:"required,max=100,min=1" json:"description"`
	IconName     string `validate:"required,max=20,min=1" json:"icon_name"`
	IconColor    string `validate:"required,max=10,min=1" json:"icon_color"`
}
