package dto

type UserCreationRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserID   string `json:"user_id" validate:"required,ascii"`
}
