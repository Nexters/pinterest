package dto

type UserCreationRequest struct {
	Password string `json:"password" validate:"required"`
	UserID   string `json:"user_id" validate:"required,ascii"`
}
