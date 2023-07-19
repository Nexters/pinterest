package database

import (
	"fmt"

	"github.com/Nexters/pinterest/interfaces/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLDialector(settings *config.Settings) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		settings.Database.User,
		settings.Database.Password,
		settings.Database.URL,
		settings.Database.Port,
		settings.Database.Name,
	)
	return mysql.Open(dsn)
}
