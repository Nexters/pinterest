package dto

type Group struct {
	Type      string `json:"type"`
	Title     string `json:"title" validate:"required"`
	Text      string `json:"text"`
	Image     string `json:"image"`
	Order     uint   `json:"order"`
	ItemCount uint   `json:"item_count" validate:"gte=0"`
	Likes     uint   `json:"likes" validate:"gte=0"`
	Link      string `json:"link"`
	UserID    uint   `json:"user_id"`
}
