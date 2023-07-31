package dto

type Film struct {
	Title         string `json:"title" validate:"required"`
	Order         uint   `json:"order"`
	PhotoCutCount uint   `json:"photo_cut_count" validate:"gte=0"`
	Likes         uint   `json:"likes" validate:"gte=0"`
	UserID        string `json:"user_id"`
}
