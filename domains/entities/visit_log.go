package entities

import "gorm.io/gorm"

type VisitLog struct {
	gorm.Model
	UserID uint
	Name   string
	Text   string
}

func (VisitLog) tableName() string {
	return "visit_logs"
}
