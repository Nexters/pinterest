package dto

type ItemCreationRequest struct {
	Title   string `json:"title" validate:"required"`
	Text    string `json:"text"`
	Link    string `json:"link"`
	Image   string `json:"image"`
	GroupID uint   `json:"group_id" validate:"required"`
}
