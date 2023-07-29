package entities

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Order         uint   `gorm:"not null"`
	PhotoCutCount uint
	Likes         uint
	UserID        string
}

func (Film) tableName() string {
	return "films"
}
