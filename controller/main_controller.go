package controller

import (
	"github.com/gofiber/fiber/v2"
)

type HealthCheck struct {
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Hello(c *fiber.Ctx) error {

	res, err := AnException()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString(res)
}

func AnException() (res string, err error) {

	return "", nil
}
