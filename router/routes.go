package router

import (
	"github.com/jojomak13/goweb-biolerplate/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Home Handler
	app.Get("/", handler.Home)

	// Users Handler
	app.Get("/users", handler.GetUsers)
	app.Post("/users", handler.CreateUser)
}