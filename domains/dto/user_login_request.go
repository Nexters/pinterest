package dto

type UserLoginRequest struct {
	UserID   string `json:"user_id" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,numeric"`
}
