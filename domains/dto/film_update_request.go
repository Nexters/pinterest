package dto

type FilmUpdateRequest struct {
	Title  string `json:"title" validate:"required"`
	FilmID string `json:"film_id" validate:"required"`
}
