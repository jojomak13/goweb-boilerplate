package router

import (
	"github.com/jojomak13/goweb-biolerplate/handler"
	"github.com/jojomak13/goweb-biolerplate/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Home Handler
	app.Get("/", handler.Home)

	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)

	// Users Handler
	app.Get("/users", middleware.Protected(), handler.GetUsers)
	app.Post("/users", handler.CreateUser)
	app.Delete("/users/:id", handler.DeleteUser)
}
