package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Password string `gorm:"not null,size:40"`
	Email    string
	UUID     string  `gorm:"not null"`
	Group    []Group `gorm:"foreignKey:UserID"`
}

func (User) tableName() string {
	return "users"
}
