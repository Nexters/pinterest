package dto

import "time"

type UserCreationResponse struct {
	Name       string
	PageUrl    string
	Email      string
	Visitors   uint
	ThemeColor string
	Text       string
	CreatedAt  time.Time
}
