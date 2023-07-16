package database

import (
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

// NewDatabase 생성자
func NewDatabase(dialector gorm.Dialector) *Database {
	config := gorm.Config{}
	db, err := gorm.Open(dialector, &config)

	if err != nil {
		log.Fatal().Err(err)
	}

	return &Database{db}
}

func (db *Database) Init() {
	db.AutoMigrate(
		&entities.Group{},
		&entities.User{},
		&entities.Item{},
		&entities.VisitLog{},
	)
}
