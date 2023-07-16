package database

import "github.com/Nexters/pinterest/interfaces/config"

func AutoMigrate(settings *config.Settings, models ...interface{}) {
	db := NewDatabase(SQLiteDialector(settings))

	db.AutoMigrate()
}
