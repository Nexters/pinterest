package dto

type UserDetailResponse struct {
	Name     string `json:"name"`
	UserID   string `json:"user_id"`
	Visitors uint   `json:"visitors"`
	Text     string `json:"text"`
	Profile  string `gorm:"profile_img" json:"profile_img"`
}
