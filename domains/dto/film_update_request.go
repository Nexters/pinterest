package dto

type FilmUpdateRequest struct {
	Title  string `json:"title" validate:"required"`
	FilmID uint   `json:"film_id" validate:"required"`
}
