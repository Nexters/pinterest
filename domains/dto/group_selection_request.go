package dto

type GroupSelectionRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
