package routes

import (
	"github.com/Leeroyakbar/bowlnow-backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, handler *handlers.UserHandler) {
	user := app.Group("/users")
	user.Post("/register", handler.Register)
	user.Post("/login", handler.Login)
}
