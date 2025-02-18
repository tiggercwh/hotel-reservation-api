package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tiggercwh/hotel-reservation-api/types"
)

func HandleListUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Bond",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
