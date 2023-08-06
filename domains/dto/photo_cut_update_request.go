package dto

type PhotoCutUpdateRequest struct {
	ID     uint   `json:"photo_cut_id" validate:"required"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Link   string `json:"link"`
	Image  string `json:"image"`
	FilmID uint   `json:"film_id" validate:"required"`
}
