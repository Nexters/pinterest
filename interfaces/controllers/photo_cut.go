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
	pc.router.Put("", pc.editPhotoCut)
	pc.router.Delete("/:photoCutId", pc.deletePhotoCut)
}

// photo_cut
// @Summary      포토컷 ID로 포토컷 정보 가져오기
// @Description  Find Photo Cut by ID
// @Tags         photo_cut
// @Accept       json
// @Produce      json
// @Param        photo_cut_id   path     uint  true  "photo_cut_id"
// @Success      200  {object}  dto.PhotoCutDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 photo_cut이 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /photo-cuts/{photo_cut_id} [get]
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

// photo_cut
// @Summary      포토컷 생성
// @Description  Create Photo Cut
// @Tags         photo_cut
// @Accept       json
// @Produce      json
// @Param        photo_cut_id   body     dto.PhotoCutCreationRequest  true  "title, text, link, image, film_id"
// @Success      201  {object}  dto.PhotoCutDetailResponse
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 film이 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /photo-cuts [post]
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
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(photoCutDto)
}

// photo_cut
// @Summary      포토컷 수정
// @Description  Edit Photo Cut
// @Tags         photo_cut
// @Accept       json
// @Produce      json
// @Param        photo_cut   body     dto.PhotoCutUpdateRequest  true  "photo_cut_id, title, text, image, link, film_id"
// @Success      200  {string}  string
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 photo_cut이 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /photo-cuts [put]
func (pc *PhotoCut) editPhotoCut(c *fiber.Ctx) error {
	dto := new(dto.PhotoCutUpdateRequest)
	err := c.BodyParser(dto)
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
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.SendString("필름 수정 성공")
}

// photo_cut
// @Summary      포토컷 삭제
// @Description  Delete Photo Cut
// @Tags         photo_cut
// @Accept       json
// @Produce      json
// @Param        photo_cut   body     dto.PhotoCutUpdateRequest  true  "photo_cut_id, title, text, image, link, film_id"
// @Success      200  {string}  string
// @failure      400              {string} string   "값을 누락하고 보냈거나, 값의 타입이 잘못된 경우"
// @failure      404              {string} string   "Conflict: 해당 id의 photo_cut이 존재하지 않는 경우"
// @failure      500  {string}   string   "Internal Server Error"
// @Router       /photo-cuts/{photo_cut_id} [delete]
func (pc *PhotoCut) deletePhotoCut(c *fiber.Ctx) error {
	photoCutIdStr := c.Params("photoCutId")
	photoCutId, err := strconv.Atoi(photoCutIdStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = pc.svc.DeletePhotoCut(c.Context(), uint(photoCutId))
	if err != nil {
		switch err.(type) {
		case *errors.NotFoundError:
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		default:
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.SendString("필름 삭제 성공")
}
