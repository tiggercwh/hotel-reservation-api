package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tiggercwh/hotel-reservation-api/api"
)

func main() {
	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("user", api.HandleListUsers)
	apiv1.Get("user/:id", api.HandleGetUser)
	app.Listen(":3000")
}

func handlerFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine!"})
}
