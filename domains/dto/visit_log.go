package dto

type VisitLog struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name" validate:"required"`
	Text   string `json:"text" validate:"required"`
}
