package middlewares

import (
	"os"
	"strings"

	"github.com/Leeroyakbar/bowlnow-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, 4001, "Missing token")
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid{
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, 4002, "Invalid token")
		}

		return c.Next()
	}
}