package controllers

import (
	"github.com/Nexters/pinterest/domains/usecases"
	"github.com/gofiber/fiber/v2"
)

type Image struct {
	router fiber.Router
	svc    *usecases.ImageService
}

func NewImageController(router fiber.Router, svc *usecases.ImageService) RouteBinder {
	return &Image{router, svc}
}

func (i *Image) Bind() {
	i.router.Get("/presigned-url", i.getPresignedUrl)
}

func (i *Image) getPresignedUrl(c *fiber.Ctx) error {
	imageDto, err := i.svc.GeneratePresignedUrl(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(imageDto)
}
