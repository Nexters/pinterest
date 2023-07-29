package controllers

import (
	"strconv"

	"github.com/Nexters/pinterest/domains/errors"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/gofiber/fiber/v2"
)

type Group struct {
	router fiber.Router
	svc    *usecases.GroupService
}

func NewGroupController(router fiber.Router, svc *usecases.GroupService) RouteBinder {
	return &Group{router, svc}
}

func (g *Group) Bind() {
	g.router.Get("/:groupId", g.getGroup)
}

func (g *Group) getGroup(c *fiber.Ctx) error {
	groupIdStr := c.Params("groupId")
	groupId, err := strconv.Atoi(groupIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	groupResponse, err := g.svc.FindByGroupId(c.Context(), uint(groupId))
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(groupResponse)
}
