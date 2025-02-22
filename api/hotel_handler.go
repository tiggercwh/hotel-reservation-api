package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tiggercwh/hotel-reservation-api/db"
	"github.com/tiggercwh/hotel-reservation-api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store *db.Store
}

func NewHotelHandler(store *db.Store) *HotelHandler {
	return &HotelHandler{
		store: store,
	}
}

func (h *HotelHandler) HandleGetHotelRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID()
	}

	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return c.JSON(rooms)
}

// func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
// 	var (
// 		params types.UpdateUserParams
// 		userID = c.Params("id")
// 	)
// 	if err := c.BodyParser(&params); err != nil {
// 		return ErrBadRequest()
// 	}
// 	filter := db.Map{"_id": userID}
// 	if err := h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
// 		return err
// 	}
// 	return c.JSON(map[string]string{"updated": userID})
// }

func (h *HotelHandler) HandleDeleteHotel(c *fiber.Ctx) error {
	hotelID := c.Params("id")
	if err := h.store.Hotel.DeleteHotel(c.Context(), hotelID); err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": hotelID})
}

func (h *HotelHandler) HandlePostHotel(c *fiber.Ctx) error {
	var params types.Hotel
	if err := c.BodyParser(&params); err != nil {
		return ErrBadRequest()
	}
	hotel, err := h.store.Hotel.InsertHotel(c.Context(), &types.Hotel{
		Name:     params.Name,
		Location: params.Location,
		Rooms:    []primitive.ObjectID{},
		Rating:   params.Rating,
	})
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(hotel)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hotel, err := h.store.Hotel.GetHotelByID(c.Context(), id)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	return c.JSON(hotel)
}

type ResourceResp struct {
	Results int `json:"results"`
	Data    any `json:"data"`
	Page    int `json:"page"`
}

type HotelQueryParams struct {
	db.Pagination
	Rating int
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var params HotelQueryParams
	if err := c.QueryParser(&params); err != nil {
		return ErrBadRequest()
	}
	filter := db.Map{
		"rating": params.Rating,
	}
	hotels, err := h.store.Hotel.GetHotels(c.Context(), filter, &params.Pagination)
	if err != nil {
		return ErrNotResourceNotFound("hotel")
	}
	resp := ResourceResp{
		Results: len(hotels),
		Data:    hotels,
		Page:    int(params.Page),
	}

	return c.JSON(resp)
}
