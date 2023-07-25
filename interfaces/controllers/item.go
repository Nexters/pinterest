package controllers

import (
	"strconv"

	"github.com/Nexters/pinterest/domains/usecases"
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
	i.router.Get("/:itemId", i.getItem)
}

func (i *Item) getItem(c *fiber.Ctx) error {
	itemIdStr := c.Params("itemId")

	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	item, err := i.svc.FindByItemId(c.Context(), uint(itemId))

	return c.JSON(item)
}
