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

	// repository
	userRepo := entities.NewUserRepository(db.DB)
	itemRepo := entities.NewItemRepository(db.DB)
	groupRepo := entities.NewGroupRepository(db.DB)

	// usecases/services
	userSvc := usecases.NewUserService(userRepo)
	itemSvc := usecases.NewItemService(itemRepo)
	groupSvc := usecases.NewGroupService(groupRepo)

	// create controllers with route Films
	root := controllers.NewRootController(app.Group("/"))
	user := controllers.NewUserController(app.Group("/user"), userSvc)
<<<<<<< HEAD
	auth := controllers.NewAuthController(app.Group("/auth"), userSvc)

	// bind routes
	controllers.BindRoutes(root, user, auth)
=======
	item := controllers.NewItemController(app.Group("/item"), itemSvc)
	group := controllers.NewGroupController(app.Group("/group"), groupSvc)

	// bind routes
	controllers.BindRoutes(root, user, item, group)
>>>>>>> 77d1ae9 (feat: add group)

	app.Listen(":8080")
}
