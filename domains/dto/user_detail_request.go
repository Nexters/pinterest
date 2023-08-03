package dto

type UserDetailRequest struct {
	UserID string `json:"user_id" validate:"required,ascii"`
}
