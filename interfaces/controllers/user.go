package controllers

import (
	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	u.router.Get("/:userId", u.getUserByID)
	u.router.Post("", u.saveUser)
	u.router.Put("", u.updateUser)
}

func (u *User) getAllUsers(c *fiber.Ctx) error {
	users, err := u.svc.FindAll(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(users)
}

// user
// @Summary      user
// @Description  Find User by ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user_id   path     string  true  "user_id"
// @Success      200  {object}  dto.UserDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /user/{user_id} [get]
func (u *User) getUserByID(c *fiber.Ctx) error {
	params := dto.UserDetailRequest{}
	err := c.ParamsParser(&params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// validate
	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userDetail, err := u.svc.FindUserByID(c.Context(), params.UserID)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(userDetail)
}

// user
// @Summary      user
// @Description  Create User
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user   body     dto.UserCreationRequest  true  "user_id, password, name(닉네임)"
// @Success      201  {object}  dto.UserCreationResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      409  {string}   string   "Conflict: 이미 id가 존재하는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /user [post]
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
		switch err {
		case gorm.ErrDuplicatedKey:
			return fiber.NewError(fiber.StatusConflict, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(userDto)
}

// user
// @Summary      user
// @Description  Create User
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user   body     dto.UserCreationRequest  true  "user_id, password, name(닉네임)"
// @Success      200  {object}  dto.UserDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /user [put]
func (u *User) updateUser(c *fiber.Ctx) error {
	userUpdateParam := dto.UserUpdateRequest{}
	err := c.BodyParser(&userUpdateParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(userUpdateParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userResponse, err := u.svc.UpdateUser(c.Context(), userUpdateParam)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(userResponse)
}
