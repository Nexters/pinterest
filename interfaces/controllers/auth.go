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

// login
// @Summary      Login
// @Description  Login user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user   body     dto.UserLoginRequest  true  "user_id, password로 로그인"
// @Success      200  {object}  dto.UserDetailResponse
// @Success      200  {object}  dto.UserDetailResponse
// @Success      200  {object}  dto.UserDetailResponse
// @failure      400              {string} string   "Key: 'UserLoginRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag"
// @failure      401  {string}   string   "Anauthorized"
// @Router       /auth [post]
func (a *Auth) login(c *fiber.Ctx) error {
	dto := new(dto.UserLoginRequest)
	err := c.BodyParser(dto)
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
