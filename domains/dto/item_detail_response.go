package dto

import "time"

type ItemDetailResponse struct {
	Title     string    `json:"title" validate:"required"`
	Text      string    `json:"text"`
	Link      string    `json:"link"`
	Image     string    `json:"image"`
	Likes     uint      `json:"likes" validate:"gte=0"`
	GroupID   uint      `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
}
