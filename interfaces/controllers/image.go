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

// image
// @Summary      presigned URL 발급
// @Description  Get Presigned URL
// @Tags         image
// @Accept       json
// @Produce      json
// @Param        filename   query   string  true  "filename"
// @Success      200  {object}  dto.ImageUploadResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /images/presigned-url [get]
func (i *Image) getPresignedUrl(c *fiber.Ctx) error {
	filename := c.Query("filename")
	if filename == "" {
		return fiber.NewError(fiber.StatusBadRequest, "파일 이름이 존재하지 않습니다.")
	}
	imageDto, err := i.svc.GeneratePresignedUrl(c.Context(), filename)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(imageDto)
}
