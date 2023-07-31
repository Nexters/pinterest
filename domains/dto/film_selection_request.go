package dto

type FilmSelectionRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
