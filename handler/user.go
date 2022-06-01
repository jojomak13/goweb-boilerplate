package handler

import (
	"net/http"

	"github.com/jojomak13/goweb-biolerplate/core"
	"github.com/jojomak13/goweb-biolerplate/model"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db := core.DB
	users := new([]model.User)

	db.Find(&users)

	return c.JSON(fiber.Map{
		"status": true,
		"users":  users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := core.DB
	user := new(model.User)

	c.BodyParser(&user)

	if errors := core.NewValidator(user); errors != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": errors,
		})
	}

	user.Password = core.HashPassword(user.Password)

	db.Create(&user)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  true,
		"message": "User created successfuly",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	db := core.DB

	db.Delete(&model.User{}, c.Params("id"))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "User deleted successfuly",
	})
}
