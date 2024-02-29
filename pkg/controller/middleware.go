package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (ctl *Controller) identityMiddleware(c *fiber.Ctx) error {
	accessToken := c.GetReqHeaders()["Authorization"]
	if len(accessToken) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(newMessage("empty token"))
	}

	headerParts := strings.Split(accessToken[0], " ")
	if len(headerParts) != 2 {
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("invalid token"))
	}
	authorId, err := parseAccessToken(headerParts[1], ctl.authConfig.JWTSigningKey)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(newMessage("invalid access token"))
	}

	c.Locals("authorId", authorId)
	return c.Next()
}
