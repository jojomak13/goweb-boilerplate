package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jojomak13/goweb-biolerplate/core"
	"github.com/jojomak13/goweb-biolerplate/model"
	request "github.com/jojomak13/goweb-biolerplate/request/auth"
)

func Login(c *fiber.Ctx) error {
	db := core.DB
	data := new(request.LoginRequest)

	c.BodyParser(&data)

	if errors := core.NewValidator(data); errors != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	user := new(model.User)
	db.Where("email", data.Email).First(&user)

	if user.ID == 0 || !core.ComparePassword(data.Password, user.Password) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid Credentials",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": true,
		"token":  core.NewJwtToken(user),
	})
}
