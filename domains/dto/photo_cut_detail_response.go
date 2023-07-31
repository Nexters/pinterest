package dto

import "time"

type PhotoCutDetailResponse struct {
	Title     string    `json:"title" validate:"required"`
	Text      string    `json:"text"`
	Link      string    `json:"link"`
	Image     string    `json:"image"`
	Likes     uint      `json:"likes" validate:"gte=0"`
	FilmID    uint      `json:"film_id"`
	CreatedAt time.Time `json:"created_at"`
}
