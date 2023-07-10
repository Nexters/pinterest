package controllers

import "github.com/gofiber/fiber/v2"

type Root struct {
}

func NewRootController() *Root {
	return &Root{}
}

func (Root) Alive(c *fiber.Ctx) error {
	return c.SendString("ok")
}
