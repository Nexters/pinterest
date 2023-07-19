package controllers

import (
	"github.com/Nexters/pinterest/domains/entities"
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
	u.router.Post("", u.saveUser)
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

func (u *User) saveUser(c *fiber.Ctx) error {
	user := &entities.User{
		Name:       "맹",
		Password:   "1234",
		Email:      "asdf123@naver.com",
		PageUrl:    "thisispageUrl",
		Visitors:   10,
		ThemeColor: "#FFFFFF",
		Text:       "너무 더워효",
	}

	err := u.svc.CreateUser(c.Context(), user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendString("저장 완료")
}
