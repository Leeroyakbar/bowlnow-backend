package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userId uuid.UUID, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}