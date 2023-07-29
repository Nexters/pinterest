package dto

type Film struct {
	Type          string `json:"type"`
	Title         string `json:"title" validate:"required"`
	Text          string `json:"text"`
	Image         string `json:"image"`
	Order         uint   `json:"order"`
	PhotoCutCount uint   `json:"photo_cut_count" validate:"gte=0"`
	Likes         uint   `json:"likes" validate:"gte=0"`
	Link          string `json:"link"`
	UserID        string `json:"user_id"`
}
