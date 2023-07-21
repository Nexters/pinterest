package dto

type UserDetailResponse struct {
	Name       string
	PageUrl    string
	Group      []Group
	VisitLog   []VisitLog
	Visitors   uint
	ThemeColor string
	Text       string
}
