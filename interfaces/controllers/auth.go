package controllers

import (
	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
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

	// validate
	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userDetail, err := a.svc.LoginUser(c.Context(), *dto)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(userDetail)
}