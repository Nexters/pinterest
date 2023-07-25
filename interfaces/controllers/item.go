package controllers

import (
	"strconv"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/errors"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	router fiber.Router
	svc    *usecases.ItemService
}

func NewItemController(router fiber.Router, isvc *usecases.ItemService) RouteBinder {
	return &Item{router, isvc}
}

func (i *Item) Bind() {
	i.router.Post("", i.saveItem)
	i.router.Get("/:itemId", i.getItem)
}

func (i *Item) getItem(c *fiber.Ctx) error {
	itemIdStr := c.Params("itemId")

	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	item, err := i.svc.FindByItemId(c.Context(), uint(itemId))
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(item)
}

func (i *Item) saveItem(c *fiber.Ctx) error {
	var itemCreationRequest dto.ItemCreationRequest
	err := c.BodyParser(&itemCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(itemCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	itemDto, err := i.svc.CreateItem(c.Context(), itemCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(itemDto)
}
