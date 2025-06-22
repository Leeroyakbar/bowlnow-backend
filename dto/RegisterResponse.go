package dto

import "github.com/google/uuid"

type RegisterResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	RoleID   uuid.UUID `json:"role_id"`
	FullName string    `json:"full_name"`
	UserName string    `json:"user_name"`
}
