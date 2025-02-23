package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tiggercwh/hotel-reservation-api/types"
)

type genericResp struct {
	Type string `json:"type"`
	Msg  string `json:"msg"`
}

func getAuthUser(c *fiber.Ctx) (*types.User, error) {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return user, nil
}
