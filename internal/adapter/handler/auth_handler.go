package handler

import (
	"backend_golang/internal/core/port"
	"backend_golang/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase port.AuthUsecase
}

func NewAuthHandler(authUsecase port.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Name == "" || req.Email == "" || req.Phone == "" || req.Password == "" {
		return response.Error(c, fiber.StatusBadRequest, "All fields are required")
	}

	user, err := h.authUsecase.Register(req.Name, req.Email, req.Phone, req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusConflict, err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "User registered successfully", fiber.Map{
		"user_id": user.ID,
		"name":    user.Name,
		"email":   user.Email,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}

	token, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Login successful", fiber.Map{
		"token": token,
	})
}
