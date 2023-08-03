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
	f.router.Put("", f.editFilm)
	f.router.Delete("/:filmId", f.deleteFilm)
}

func (f *Film) getFilm(c *fiber.Ctx) error {
	filmIdStr := c.Params("filmId")
	filmId, err := strconv.Atoi(filmIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
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

func (f *Film) saveFilm(c *fiber.Ctx) error {
	dto := new(dto.FilmCreationRequest)
	err := c.BodyParser(&dto)
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
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(filmDto)
}

func (f *Film) getAllFilms(c *fiber.Ctx) error {
	dto := new(dto.FilmSelectionRequest)
	err := c.BodyParser(&dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	films, err := f.svc.FindAllFilms(c.Context(), dto.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(films)
}

func (f *Film) editFilm(c *fiber.Ctx) error {
	dto := new(dto.FilmUpdateRequest)
	err := c.BodyParser(&dto)
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
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("필름 수정 성공")
}

func (f *Film) deleteFilm(c *fiber.Ctx) error {
	filmIdStr := c.Params("filmId")
	filmId, err := strconv.Atoi(filmIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = f.svc.DeleteFilm(c.Context(), uint(filmId))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("필름 삭제 성공")
}
