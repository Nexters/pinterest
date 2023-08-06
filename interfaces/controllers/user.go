package controllers

import (
	"strconv"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User struct {
	router  fiber.Router
	userSvc *usecases.UserService
	logSvc  *usecases.VisitLogService
}

func NewUserController(router fiber.Router, userSvc *usecases.UserService, logSvc *usecases.VisitLogService) RouteBinder {
	return &User{router, userSvc, logSvc}
}

func (u *User) Bind() {
	u.router.Get("", u.getAllUsers)
	u.router.Get("/:userId", u.getUserByID)
	u.router.Get("/:userId/visit-logs", u.getAllVisitLogs)
	u.router.Post("/:userId/visit-logs", u.saveVisitLog)
	u.router.Delete("/:userId/visit-logs/:logId", u.deleteVisitLog)
	u.router.Post("", u.saveUser)
	u.router.Put("", u.updateUser)
}

func (u *User) getAllUsers(c *fiber.Ctx) error {
	users, err := u.userSvc.FindAll(c.Context())
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

	userDetail, err := u.userSvc.FindUserByID(c.Context(), params.UserID)

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

	userDto, err := u.userSvc.CreateUser(c.Context(), userCreationRequest)
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
// @Success      200  {object}  dto.UserUpdateResponse
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

	userResponse, err := u.userSvc.UpdateUser(c.Context(), userUpdateParam)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(userResponse)
}

// user
// @Summary      user
// @Description  Find all visit logs from given user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user_id   path     string  true  "user_id"
// @Success      200  {object}  dto.UserDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /user/{user_id}/visit-logs [get]
func (u *User) getAllVisitLogs(c *fiber.Ctx) error {
	userID := c.Params("userId")
	logs, err := u.logSvc.FindAll(c.Context(), userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(logs)
}

// user
// @Summary      user
// @Description  Create User's visit log
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user   body     dto.VisitLogCreationRequest  true  "visit log 생성"
// @Success      201  {object}  dto.VisitLogCreationResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      409  {string}   string   "Conflict: 이미 id가 존재하는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /user/{user_id}/visit-logs [post]
func (u *User) saveVisitLog(c *fiber.Ctx) error {
	userID := c.Params("userId")
	dto := dto.VisitLogCreationRequest{
		UserID: userID,
	}
	err := c.BodyParser(&dto)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	log, err := u.logSvc.Create(c.Context(), dto)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(log)
}

func (u *User) deleteVisitLog(c *fiber.Ctx) error {
	logID := c.Params("logId")
	i, err := strconv.ParseUint(logID, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = u.logSvc.Delete(c.Context(), uint(i))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("ok")
}
