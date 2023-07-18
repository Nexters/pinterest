package main

import (
	"github.com/Nexters/pinterest/interfaces/config"
	"github.com/Nexters/pinterest/interfaces/controllers"
	"github.com/Nexters/pinterest/interfaces/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	settings := config.NewSettings()

	// create controllers with route groups
	root := controllers.NewRootController(app.Group("/"))
	user := controllers.NewUserController(app.Group("/user"))

	// bind routes
	controllers.BindRoutes(root, user)

	// Database
	db := database.NewDatabase(database.MySQLDialector(settings))
	db.Init()

	app.Listen(":8080")
}
