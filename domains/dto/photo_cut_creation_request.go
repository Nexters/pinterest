package dto

type PhotoCutCreationRequest struct {
	Title  string `json:"title" validate:"required"`
	Text   string `json:"text"`
	Link   string `json:"link"`
	Image  string `json:"image"`
	FilmID uint   `json:"film_id" validate:"required"`
}
