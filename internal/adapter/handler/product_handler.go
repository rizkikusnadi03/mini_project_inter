package handler

import (
	"strconv"

	"backend_golang/internal/core/port"
	"backend_golang/pkg/response"
	"backend_golang/pkg/upload"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUsecase port.ProductUsecase
}

func NewProductHandler(productUsecase port.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase}
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64) // JWT map claims parse numbers as float64

	categoryID, _ := strconv.ParseUint(c.FormValue("category_id"), 10, 32)
	name := c.FormValue("name")
	description := c.FormValue("description")
	price, _ := strconv.ParseFloat(c.FormValue("price"), 64)
	stock, _ := strconv.Atoi(c.FormValue("stock"))

	if name == "" || categoryID == 0 || price <= 0 || stock < 0 {
		return response.Error(c, fiber.StatusBadRequest, "Invalid input fields")
	}

	imagePath, err := upload.SaveImage(c, "image", "products")
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to upload image")
	}

	product, err := h.productUsecase.CreateProduct(uint(userID), uint(categoryID), name, description, price, stock, imagePath)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, total, err := h.productUsecase.GetProducts(page, limit, search)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch products")
	}

	return response.Success(c, fiber.StatusOK, "Success", fiber.Map{
		"products": products,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}
