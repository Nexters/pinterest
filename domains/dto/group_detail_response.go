package dto

type GroupDetailResponse struct {
	GroupID   uint                 `json:"group_id"`
	Title     string               `json:"title"`
	Order     uint                 `json:"order"`
	ItemCount uint                 `json:"item_count"`
	Likes     uint                 `json:"likes"`
	UserID    string               `json:"user_id"`
	Items     []ItemDetailResponse `json:"items"`
}
