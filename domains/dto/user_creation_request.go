package dto

type UserCreationRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	PageUrl  string `json:"page_url"`
	UserID   string `json:"user_id" validate:"required"`
}
