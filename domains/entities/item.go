package entities

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Text    string
	Link    string
	Image   string
	Likes   uint
	GroupID uint
}

func (Item) tableName() string {
	return "items"
}
