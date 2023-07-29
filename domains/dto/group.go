package dto

type Group struct {
	Title     string `json:"title" validate:"required"`
	Order     uint   `json:"order"`
	ItemCount uint   `json:"item_count" validate:"gte=0"`
	Likes     uint   `json:"likes" validate:"gte=0"`
	UserID    string `json:"user_id"`
}
