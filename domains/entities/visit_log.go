package entities

import "gorm.io/gorm"

type VisitLog struct {
	gorm.Model
	UserID uint
	name   string
	text   string
}

func (VisitLog) tableName() string {
	return "visit_logs"
}
