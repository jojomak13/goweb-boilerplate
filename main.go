package main

import (
	"log"

	"github.com/jojomak13/goweb-biolerplate/core"
	"github.com/jojomak13/goweb-biolerplate/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	core.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}