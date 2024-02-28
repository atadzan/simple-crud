package controller

import (
	"github.com/atadzan/simple-crud/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func (ctl *Controller) register(c *fiber.Ctx) error {
	var input models.AuthParams
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusOK)
	}
	return nil
}

func (ctl *Controller) signIn(c *fiber.Ctx) error {

	return nil
}
