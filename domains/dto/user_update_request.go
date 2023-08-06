package dto

import "github.com/Nexters/pinterest/domains/entities"

type UserUpdateRequest struct {
	Name       string `json:"name"`
	Password   string `json:"password" validate:"numeric"`
	Email      string
	ThemeColor string `json:"theme_color"`
	Text       string
	Profile    string `json:"profile_img"`
	UserID     string `json:"user_id" validate:"required,ascii"`
}

func (u UserUpdateRequest) ToEntity() entities.User {
	return entities.User{
		Name:       u.Name,
		Password:   u.Password,
		Email:      u.Email,
		ThemeColor: u.ThemeColor,
		Text:       u.Text,
		Profile:    u.Profile,
		ID:         u.UserID,
	}
}
