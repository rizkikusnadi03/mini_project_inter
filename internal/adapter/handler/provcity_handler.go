package handler

import (
	"backend_golang/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type ProvCityHandler struct{}

func NewProvCityHandler() *ProvCityHandler {
	return &ProvCityHandler{}
}

func (h *ProvCityHandler) GetProvinces(c *fiber.Ctx) error {
	provinces := []fiber.Map{
		{"id": "prov-1", "name": "Jawa Barat"},
		{"id": "prov-2", "name": "DKI Jakarta"},
		{"id": "prov-3", "name": "Banten"},
	}
	return response.Success(c, fiber.StatusOK, "Success", provinces)
}

func (h *ProvCityHandler) GetCities(c *fiber.Ctx) error {
	provID := c.Params("prov_id")

	var cities []fiber.Map

	switch provID {
	case "prov-1":
		cities = []fiber.Map{
			{"id": "city-1", "name": "Bandung"},
			{"id": "city-2", "name": "Cimahi"},
		}
	case "prov-2":
		cities = []fiber.Map{
			{"id": "city-3", "name": "Jakarta Selatan"},
			{"id": "city-4", "name": "Jakarta Pusat"},
		}
	case "prov-3":
		cities = []fiber.Map{
			{"id": "city-5", "name": "Tangerang"},
			{"id": "city-6", "name": "Serang"},
		}
	default:
		return response.Error(c, fiber.StatusNotFound, "Province not found")
	}

	return response.Success(c, fiber.StatusOK, "Success", cities)
}
