package dto

import "github.com/google/uuid"

type LoginResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name"`
	FullName string    `json:"full_name"`
	RoleName string    `json:"role_name"`
	Token    string    `json:"token"`
}
