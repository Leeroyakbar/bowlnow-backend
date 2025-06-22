package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	RoleID       uuid.UUID `gorm:"type:uuid" json:"role_id"`
	FullName     string    `gorm:"type:varchar(100)" json:"full_name"`
	UserName     string    `gorm:"type:varchar(50);unique" json:"user_name"`
	Password     string    `gorm:"type:varchar(255)" json:"password"`
	Image        string    `gorm:"type:varchar(100)"`
	GuestFlag    int       `gorm:"type:integer" json:"guest_flag"`
	DeleteStatus int       `gorm:"type:integer" json:"delete_status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Role Role
}

func (User) TableName() string {
	return "mst_users"
}
