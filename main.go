package main

import (
	"fmt"
	"log"

	"github.com/jojomak13/goweb-biolerplate/config"
	"github.com/jojomak13/goweb-biolerplate/core"
	"github.com/jojomak13/goweb-biolerplate/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config := config.LoadConfigration()

	core.ConnectDB(config)

	router.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.Port)))
}