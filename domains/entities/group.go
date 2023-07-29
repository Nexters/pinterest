package entities

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Type      string `gorm:"not null"`
	Title     string `gorm:"not null"`
	Text      string
	Image     string `gorm:"not null"`
	Order     uint   `gorm:"not null"`
	ItemCount uint
	Likes     uint
	Link      string
	UserID    uint
	Items     []Item
}

func (Group) tableName() string {
	return "groups"
}
