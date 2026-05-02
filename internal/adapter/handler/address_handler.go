package handler

import (
	"backend_golang/internal/core/port"
	"backend_golang/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	addressUsecase port.AddressUsecase
}

func NewAddressHandler(addressUsecase port.AddressUsecase) *AddressHandler {
	return &AddressHandler{addressUsecase}
}

func (h *AddressHandler) GetMyAddresses(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	addresses, err := h.addressUsecase.GetAddressesByUserID(uint(userID))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get addresses")
	}

	return response.Success(c, fiber.StatusOK, "Success", addresses)
}

func (h *AddressHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	type CreateRequest struct {
		Title          string `json:"title"`
		AddressDetails string `json:"address_details"`
		ProvID         string `json:"prov_id"`
		CityID         string `json:"city_id"`
		IsPrimary      bool   `json:"is_primary"`
	}

	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	address, err := h.addressUsecase.CreateAddress(uint(userID), req.Title, req.AddressDetails, req.ProvID, req.CityID, req.IsPrimary)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create address")
	}

	return response.Success(c, fiber.StatusCreated, "Address created successfully", address)
}
