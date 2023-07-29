package dto

import "time"

type UserCreationResponse struct {
	Name string
	string
	Email      string
	Visitors   uint
	ThemeColor string
	Text       string
	CreatedAt  time.Time
}
