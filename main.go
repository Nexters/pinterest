package main

import (
	_ "github.com/Nexters/pinterest/docs"
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/Nexters/pinterest/interfaces/config"
	"github.com/Nexters/pinterest/interfaces/controllers"
	"github.com/Nexters/pinterest/interfaces/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler

	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title grafi API
// @version 1.0
// @description grafi API
// @contact.email chaewonkong@gmail.com
// @host api.grafi.cc
// @BasePath /
func main() {
	app := fiber.New()
	settings := config.NewSettings()

	// Initialize default config
	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Database
	db := database.NewDatabase(database.MySQLDialector(settings))

	// repository
	userRepo := entities.NewUserRepository(db.DB)
	photoCutRepo := entities.NewPhotoCutRepository(db.DB)
	filmRepo := entities.NewFilmRepository(db.DB)
	logRepo := entities.NewVisitLogRepository(db.DB)

	// usecases/services
	userSvc := usecases.NewUserService(userRepo)
	photoCutSvc := usecases.NewPhotoCutService(photoCutRepo, filmRepo)
	filmSvc := usecases.NewFilmService(filmRepo, userRepo)
	logSvc := usecases.NewVisitLogService(logRepo)

	imageSvc := usecases.NewImageService()

	// create controllers with route Films
	root := controllers.NewRootController(app.Group("/"))
	user := controllers.NewUserController(app.Group("/users"), userSvc, logSvc)
	auth := controllers.NewAuthController(app.Group("/auth"), userSvc)
	photo_cut := controllers.NewPhotoCutController(app.Group("/photo-cuts"), photoCutSvc)
	film := controllers.NewFilmController(app.Group("/films"), filmSvc)
	image := controllers.NewImageController(app.Group("/images"), imageSvc)

	// bind routes
	controllers.BindRoutes(root, user, auth, photo_cut, film, image)

	app.Listen(":8080")
}
