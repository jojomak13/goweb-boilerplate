package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome My App",
		"version": "v1.0",
	})
}
