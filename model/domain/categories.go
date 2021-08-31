package domain

import (
	"time"
)

type Category struct {
	Id           uint8     `json:"id"`
	UserId       uint8     `json:"user_id"`
	NameCategory string    `json:"name_category"`
	Description  string    `json:"description"`
	IconName     string    `json:"icon_name"`
	IconColor    string    `json:"icon_color"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
