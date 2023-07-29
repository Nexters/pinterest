package dto

type GroupCreationRequest struct {
	Title  string `json:"title" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}
