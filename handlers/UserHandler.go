package handlers

import (
	"github.com/Leeroyakbar/bowlnow-backend/dto"
	"github.com/Leeroyakbar/bowlnow-backend/services"
	"github.com/Leeroyakbar/bowlnow-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, 4001, "Invalid form data")
	}

	fileHeader := form.File["image"]
	if len(fileHeader) == 0 {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, 4002, "Image is required")
	}

	user, err := h.userService.RegisterFromForm(c)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, 4003, err.Error())
	}

	return utils.SuccessResponse(c, user, "User registered successfully")
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, 4001, "Invalid request body")
	}

	res, err := h.userService.Login(req.UserName, req.Password) 
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, 4002, err.Error())
	}

	return utils.SuccessResponse(c, res, "Login successful")
}