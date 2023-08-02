package controllers

import (
	"strconv"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/errors"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type PhotoCut struct {
	router fiber.Router
	svc    *usecases.PhotoCutService
}

func NewPhotoCutController(router fiber.Router, svc *usecases.PhotoCutService) RouteBinder {
	return &PhotoCut{router, svc}
}

func (pc *PhotoCut) Bind() {
	pc.router.Post("", pc.savePhotoCut)
	pc.router.Get("/:photoCutId", pc.getPhotoCut)
}

func (pc *PhotoCut) getPhotoCut(c *fiber.Ctx) error {
	photoCutIdStr := c.Params("photoCutId")

	photoCutId, err := strconv.Atoi(photoCutIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	photoCut, err := pc.svc.FindByPhotoCutId(c.Context(), uint(photoCutId))
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(photoCut)
}

func (pc *PhotoCut) savePhotoCut(c *fiber.Ctx) error {
	var photoCutCreationRequest dto.PhotoCutCreationRequest
	err := c.BodyParser(&photoCutCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(photoCutCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	photoCutDto, err := pc.svc.CreatePhotoCut(c.Context(), photoCutCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(photoCutDto)
}

func (pc *PhotoCut) editPhotoCut(c *fiber.Ctx) error {
	dto := new(dto.PhotoCutUpdateRequest)
	err := c.BodyParser(&dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(dto)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = pc.svc.UpdatePhotoCut(c.Context(), *dto)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("필름 수정 성공")
}

func (pc *PhotoCut) deletePhotoCut(c *fiber.Ctx) error {
	photoCutIdStr := c.Params("photoCutId")
	photoCutId, err := strconv.Atoi(photoCutIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = pc.svc.DeletePhotoCut(c.Context(), uint(photoCutId))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("필름 삭제 성공")
}
