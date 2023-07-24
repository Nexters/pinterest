package dto

type UserDetailResponse struct {
	Name       string     `json:"name"`
	PageUrl    string     `json:"page_url"`
	Group      []Group    `json:"groups"`
	VisitLog   []VisitLog `json:"visit_logs"`
	Visitors   uint       `json:"visitors" validate:"gte=0"`
	ThemeColor string     `json:"theme_color"`
	Text       string     `json:"text"`
}
