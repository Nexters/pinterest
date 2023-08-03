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
	// 생성된 Presigned URL을 클라이언트에게 전달
	return c.JSON(imageDto)
}
