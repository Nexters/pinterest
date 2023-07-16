package database

import (
	"fmt"

	"github.com/Nexters/pinterest/interfaces/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLiteDialector(settings *config.Settings) gorm.Dialector {
	dsn := fmt.Sprintf("%s.db", settings.Database.Name)

	return sqlite.Open(dsn)
}
