package utils

import (
	"github.com/Leeroyakbar/bowlnow-backend/dto"
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse[T any](c *fiber.Ctx, data T, message string) error {
	res := dto.APIResponse[T]{
		Success:      true,
		ErrorCode:    0,
		ErrorMessage: "",
		Data:         data,
		Message:      message,
	}

	return c.JSON(res)
}

func ErrorResponse(c *fiber.Ctx, statusCode int, errorCode int, errorMsg string) error {
	res := dto.APIResponse[any]{
		Success:      false,
		ErrorCode:    errorCode,
		ErrorMessage: errorMsg,
		Data:         nil,
		Message:      "Error occurred",
	}

	return c.Status(statusCode).JSON(res)
}
