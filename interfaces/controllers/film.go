package controllers

import (
	"strconv"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/errors"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Film struct {
	router fiber.Router
	svc    *usecases.FilmService
}

func NewFilmController(router fiber.Router, svc *usecases.FilmService) RouteBinder {
	return &Film{router, svc}
}

func (f *Film) Bind() {
	f.router.Get("/:filmId", f.getFilm)
	f.router.Post("", f.saveFilm)
	f.router.Get("", f.getAllFilms)
	f.router.Put("", f.editFilm)
	f.router.Delete("/:filmId", f.deleteFilm)
}

// film
// @Summary      필름 ID로 필름 정보 가져오기
// @Description  Find Film by ID
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        film_id   path     uint  true  "film_id"
// @Success      200  {object}  dto.FilmDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 film이 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /films/{film_id} [get]
func (f *Film) getFilm(c *fiber.Ctx) error {
	filmIdStr := c.Params("filmId")
	filmId, err := strconv.Atoi(filmIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	filmResponse, err := f.svc.FindByFilmId(c.Context(), uint(filmId))
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(filmResponse)
}

// film
// @Summary      필름 생성
// @Description  Create Film
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        film   body     dto.FilmCreationRequest  true  "user_id, title"
// @Success      201  {object}  dto.FilmDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 user가 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /films [post]
func (f *Film) saveFilm(c *fiber.Ctx) error {
	dto := new(dto.FilmCreationRequest)
	err := c.BodyParser(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	filmDto, err := f.svc.CreateFilm(c.Context(), *dto)
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(filmDto)
}

// film
// @Summary      회원의 모든 필름 가져오기
// @Description  Find All Films
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        user_id   query   string  true  "user_id"
// @Success      200  {object}  []dto.FilmDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /films [get]
func (f *Film) getAllFilms(c *fiber.Ctx) error {
	userId := c.Query("user_id")
	if userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "필수 값이 없습니다.")
	}

	films, err := f.svc.FindAllFilms(c.Context(), userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(films)
}

// film
// @Summary      필름 수정
// @Description  Edit Film
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        film   body     dto.FilmUpdateRequest  true  "film_id, title"
// @Success      200  {string}  string
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 film이 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /films [put]
func (f *Film) editFilm(c *fiber.Ctx) error {
	dto := new(dto.FilmUpdateRequest)
	err := c.BodyParser(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = f.svc.UpdateFilm(c.Context(), *dto)
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.SendString("필름 수정 성공")
}

// film
// @Summary      필름 삭제
// @Description  Delete Film
// @Tags         film
// @Accept       json
// @Produce      json
// @Param        film_id   path     uint  true  "film_id"
// @Success      200  {string}  string
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 film이 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /films/{film_id} [delete]
func (f *Film) deleteFilm(c *fiber.Ctx) error {
	filmIdStr := c.Params("filmId")
	filmId, err := strconv.Atoi(filmIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = f.svc.DeleteFilm(c.Context(), uint(filmId))
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.SendString("필름 삭제 성공")
}
