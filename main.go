package main

import (
	"github.com/Nexters/pinterest/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hi")
	})

	h := controller.NewHealthCheck()
	app.Get("/", h.Hello)

	app.Listen(":8080")
}
