package routes

import (
	"gofiber/users"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", users.GetAllUser)
}
