package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	RoleID       uuid.UUID `gorm:"type:uuid" json:"role_id"`
	FullName     string    `gorm:"type:varchar(100)" json:"full_name"`
	UserName     string    `gorm:"type:varchar(50);unique" json:"user_name"`
	Password     string    `gorm:"type:varchar(255)" json:"password"`
	Image        string    `gorm:"type:varchar(100)"`
	GuestFlag    int       `gorm:"type:integer" json:"guest_flag"`
	DeleteStatus string    `gorm:"type:integer" json:"delete_status"`
	CreatedAt    time.Time `gorm:"type:uuid" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:uuid" json:"updated_at"`

	// relasi
	Role Role
}

func (User) TableName() string {
	return "mst_users"
}
