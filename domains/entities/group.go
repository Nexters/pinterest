package entities

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Order     uint   `gorm:"not null"`
	ItemCount uint
	Likes     uint
	UserID    string
	Items     []Item
}

func (Group) tableName() string {
	return "groups"
}
