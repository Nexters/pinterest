package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         string `gorm:"primaryKey;not null"`
	Name       string `gorm:"not null"`
	Password   string `gorm:"not null,size:40"`
	Email      string
	Visitors   uint
	ThemeColor string
	Text       string
	Profile    string `json:"profile_img"`
}

func (User) tableName() string {
	return "users"
}
