package handler

import (
	"strconv"

	"backend_golang/internal/core/port"
	"backend_golang/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryUsecase port.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase port.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase}
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	type CreateRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var req CreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Name == "" {
		return response.Error(c, fiber.StatusBadRequest, "Category name is required")
	}

	category, err := h.categoryUsecase.CreateCategory(req.Name, req.Description)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "Category created successfully", category)
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	type UpdateRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	category, err := h.categoryUsecase.UpdateCategory(uint(id), req.Name, req.Description)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Category not found or update failed")
	}

	return response.Success(c, fiber.StatusOK, "Category updated successfully", category)
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid category ID")
	}

	err = h.categoryUsecase.DeleteCategory(uint(id))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete category")
	}

	return response.Success(c, fiber.StatusOK, "Category deleted successfully", nil)
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	categories, err := h.categoryUsecase.GetCategories()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get categories")
	}

	return response.Success(c, fiber.StatusOK, "Success", categories)
}
