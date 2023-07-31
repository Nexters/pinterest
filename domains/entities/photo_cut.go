package entities

import "gorm.io/gorm"

type PhotoCut struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Text   string
	Link   string
	Image  string
	Likes  uint
	FilmID uint
}

func (PhotoCut) tableName() string {
	return "photo_cuts"
}
