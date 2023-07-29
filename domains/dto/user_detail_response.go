package dto

type UserDetailResponse struct {
	Name     string `json:"name"`
	Visitors uint   `json:"visitors" validate:"gte=0"`
	Text     string `json:"text"`
	Profile  string `gorm:"profile_img" json:"profile_img"`
}
