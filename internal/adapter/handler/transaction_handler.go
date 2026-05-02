package handler

import (
	"backend_golang/internal/core/port"
	"backend_golang/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	transactionUsecase port.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase port.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{transactionUsecase}
}

func (h *TransactionHandler) Checkout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	type CheckoutRequest struct {
		ProductID     uint   `json:"product_id"`
		Quantity      int    `json:"quantity"`
		PaymentMethod string `json:"payment_method"`
	}

	var req CheckoutRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	transaction, err := h.transactionUsecase.Checkout(uint(userID), req.ProductID, req.Quantity, req.PaymentMethod)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Checkout successful", transaction)
}

func (h *TransactionHandler) GetMyTransactions(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)

	transactions, err := h.transactionUsecase.GetUserTransactions(uint(userID))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get transactions")
	}

	return response.Success(c, fiber.StatusOK, "Success", transactions)
}
