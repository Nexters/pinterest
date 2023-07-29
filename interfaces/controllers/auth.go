package controllers

import (
	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	router fiber.Router
	svc    *usecases.UserService
}

func NewAuthController(router fiber.Router, svc *usecases.UserService) RouteBinder {
	return &Auth{router, svc}
}

func (a *Auth) Bind() {
	a.router.Post("", a.login)
}

func (a *Auth) login(c *fiber.Ctx) error {
	dto := new(dto.UserLoginRequest)
	err := c.BodyParser(&dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = a.svc.LoginUser(c.Context(), *dto)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}
	return c.SendString("ok")
}
