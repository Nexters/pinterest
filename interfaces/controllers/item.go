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

func NewPhotoCutController(router fiber.Router, isvc *usecases.PhotoCutService) RouteBinder {
	return &PhotoCut{router, isvc}
}

func (i *PhotoCut) Bind() {
	i.router.Post("", i.savePhotoCut)
	i.router.Get("/:photoCutId", i.getPhotoCut)
}

func (i *PhotoCut) getPhotoCut(c *fiber.Ctx) error {
	photoCutIdStr := c.Params("photoCutId")

	photoCutId, err := strconv.Atoi(photoCutIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	photoCut, err := i.svc.FindByPhotoCutId(c.Context(), uint(photoCutId))
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

func (i *PhotoCut) savePhotoCut(c *fiber.Ctx) error {
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

	photoCutDto, err := i.svc.CreatePhotoCut(c.Context(), photoCutCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(photoCutDto)
}
