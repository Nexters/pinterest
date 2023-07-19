package main

import (
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/Nexters/pinterest/interfaces/config"
	"github.com/Nexters/pinterest/interfaces/controllers"
	"github.com/Nexters/pinterest/interfaces/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	settings := config.NewSettings()

	// Database
	db := database.NewDatabase(database.MySQLDialector(settings))

	db.Init()

	// repository
	userRepo := entities.NewUserRepository(db.DB)

	// usecases/services
	userSvc := usecases.NewUserService(userRepo)

	// create controllers with route groups
	root := controllers.NewRootController(app.Group("/"))
	user := controllers.NewUserController(app.Group("/user"), userSvc)

	// bind routes
	controllers.BindRoutes(root, user)

	app.Listen(":8080")
}
