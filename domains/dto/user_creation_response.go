package dto

import "time"

type UserCreationResponse struct {
	Name       string    `json:"name"`
	UserID     string    `json:"user_id"`
	Visitors   uint      `json:"visitors"`
	ThemeColor string    `json:"theme_colors"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
