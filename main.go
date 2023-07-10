package main

import (
	"github.com/Nexters/pinterest/interfaces/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// TODO: routing 분리
	rc := controllers.NewRootController()
	app.Get("/", rc.Alive)

	app.Listen(":8080")
}
