package web

type CategoryCreateRequest struct {
	UserId       uint8  `json:"user_id"`
	NameCategory string `json:"name_category"`
	Description  string `json:"description"`
	IconName     string `json:"icon_name"`
	IconColor    string `json:"icon_color"`
}
