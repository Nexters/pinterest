package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID     string `gorm:"not null"`
	Name       string `gorm:"not null"`
	Password   string `gorm:"not null,size:40"`
	Email      string
	PageUrl    string     `gorm:"not null"`
	Group      []Group    `gorm:"foreignKey:UserID" json:"groups"`
	VisitLog   []VisitLog `gorm:"foreignKey:UserID" json:"visit_logs"`
	Visitors   uint
	ThemeColor string
	Text       string
}

func (User) tableName() string {
	return "users"
}
