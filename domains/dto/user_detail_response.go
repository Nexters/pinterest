package dto

import "github.com/Nexters/pinterest/domains/entities"

type UserDetailResponse struct {
	Name     string  `json:"name"`
	UserID   string  `json:"user_id"`
	Visitors uint    `json:"visitors"`
	Text     *string `json:"text"`
	Profile  *string `json:"profile_img"`
}

func (u UserDetailResponse) FromEntity(user entities.User) UserDetailResponse {
	u.Name = user.Name
	u.UserID = user.ID
	u.Visitors = user.Visitors
	u.Text = ConvertEmptyStringToNil(&user.Text)
	u.Profile = ConvertEmptyStringToNil(&user.Profile)

	return UserDetailResponse{
		Name:     user.Name,
		UserID:   user.ID,
		Visitors: user.Visitors,
		Text:     ConvertEmptyStringToNil(&user.Text),
		Profile:  ConvertEmptyStringToNil(&user.Profile),
	}
}

func ConvertEmptyStringToNil(s *string) *string {
	if *s == "" {
		return nil
	}

	return s
}
