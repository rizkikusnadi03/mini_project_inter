package handler

import (
	"strconv"

	"backend_golang/internal/core/port"
	"backend_golang/pkg/response"
	"backend_golang/pkg/upload"
	"github.com/gofiber/fiber/v2"
)

type StoreHandler struct {
	storeUsecase port.StoreUsecase
}

func NewStoreHandler(storeUsecase port.StoreUsecase) *StoreHandler {
	return &StoreHandler{storeUsecase}
}

func (h *StoreHandler) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)
	storeIDStr := c.Params("id_toko")
	storeID, err := strconv.ParseUint(storeIDStr, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid store ID")
	}

	name := c.FormValue("name")
	description := c.FormValue("description")

	imagePath, err := upload.SaveImage(c, "image", "stores")
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to upload image")
	}

	store, err := h.storeUsecase.UpdateStore(uint(userID), uint(storeID), name, description, imagePath)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Store updated successfully", store)
}
