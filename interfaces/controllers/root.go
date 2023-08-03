package controllers

import "github.com/gofiber/fiber/v2"

type Root struct {
	router fiber.Router
}

func NewRootController(router fiber.Router) RouteBinder {
	return &Root{router}
}

// login
// @Summary      healthcheck
// @Description  healthcheck
// @Tags         default
// @Accept       json
// @Produce      json
// @Success      200  {string}  "ok"
// @Router       / [get]
func (r *Root) Bind() {
	r.router.Get("/", r.alive)
}

func (r *Root) alive(c *fiber.Ctx) error {
	return c.SendString("ok")
}
