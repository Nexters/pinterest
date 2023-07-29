package controllers

import (
	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
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
	// u.router.Get("/:userId", u.getUser)
	u.router.Post("", u.saveUser)
}

func (u *User) getAllUsers(c *fiber.Ctx) error {
	users, err := u.svc.FindAll(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(users)
}

func (u *User) saveUser(c *fiber.Ctx) error {
	var userCreationRequest dto.UserCreationRequest
	err := c.BodyParser(&userCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// UserCreationRequest 검증
	validate := validator.New()
	err = validate.Struct(userCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userDto, err := u.svc.CreateUser(c.Context(), userCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(userDto)
}
