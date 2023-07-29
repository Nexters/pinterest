package controllers

import (
	"strconv"

	"github.com/Nexters/pinterest/domains/dto"
	"github.com/Nexters/pinterest/domains/errors"
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/go-playground/validator"
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
	g.router.Post("", g.saveGroup)
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

func (g *Group) saveGroup(c *fiber.Ctx) error {
	var groupCreationRequest dto.GroupCreationRequest
	err := c.BodyParser(&groupCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(groupCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	itemDto, err := g.svc.CreateGroup(c.Context(), groupCreationRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(itemDto)
}
