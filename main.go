package main

import (
	"gofiber/config"
	"gofiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()
	app := fiber.New()
	routes.RouteInit(app)
	app.Listen(":3000")
}
