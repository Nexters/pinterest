package controllers

import (
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	router fiber.Router
	svc    *usecases.UserService
}

func NewUserController(router fiber.Router, svc *usecases.UserService) RouteBinder {
	return &User{router, svc}
}

func (u *User) Bind() {
	u.router.Get("", u.getAllUsers)
	u.router.Get("/:userId", u.getUser)
}

func (u *User) getAllUsers(c *fiber.Ctx) error {
	users, err := u.svc.FindAll(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(users)
}

func (u *User) getUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	return c.SendString(userId)
}
