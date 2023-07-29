package dto

type GroupDetailResponse struct {
	Type      string               `json:"type"`
	Title     string               `json:"title"`
	Text      string               `json:"text"`
	Image     string               `json:"image"`
	Order     uint                 `json:"order"`
	ItemCount uint                 `json:"item_count"`
	Likes     uint                 `json:"likes"`
	Link      string               `json:"link"`
	UserID    uint                 `json:"user_id"`
	Items     []ItemDetailResponse `json:"items"`
}
