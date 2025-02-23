package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tiggercwh/hotel-reservation-api/db"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store db.BookingStore
}

func NewBookingHandler(booking_store db.BookingStore) *BookingHandler {
	return &BookingHandler{
		store: booking_store,
	}
}

func (h *BookingHandler) HandleGetBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.store.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil || booking.UserID != user.ID {
		return ErrUnAuthorized()
	}
	return c.JSON(booking)
}

func (h *BookingHandler) HandleCancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")

	booking, err := h.store.GetBookingByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("booking")
	}
	user, err := getAuthUser(c)
	if err != nil {
		return ErrUnAuthorized()
	}
	if booking.UserID != user.ID {
		return ErrUnAuthorized()
	}

	err = h.store.UpdateBooking(c.Context(), id, bson.M{"canceled": true})
	if err != nil {
		return err
	}
	return c.JSON(genericResp{Type: "msg", Msg: "updated"})
}
