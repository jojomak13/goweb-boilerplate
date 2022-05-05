package handler

import (
	"net/http"

	"github.com/jojomak13/goweb-biolerplate/core"
	"github.com/jojomak13/goweb-biolerplate/model"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		"message": "Users Request",
	})
}

func CreateUser(c *fiber.Ctx) error{
	db := core.DB
	user := new(model.User)

	c.BodyParser(&user)

	if errors := core.NewValidator(user); errors != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Status": "Bad Request",
			"errors": errors,
		})
	}

	db.Create(&user)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created",
	})
}