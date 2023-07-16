package controllers

import "github.com/gofiber/fiber/v2"

type User struct {
	router fiber.Router
}

func NewUserController(router fiber.Router) RouteBinder {
	return &User{router}
}

func (u *User) Bind() {
	u.router.Get("/:userId", u.getUser)
}

func (u *User) getUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	return c.SendString(userId)
}
