package models

import "github.com/google/uuid"

type Role struct {
	RoleID       uuid.UUID `gorm:"type:uuid;primary_key" json:"role_id"`
	RoleName     string    `gorm:"type:varchar(50);unique" json:"role_name"`
	DeleteStatus string    `gorm:"type:int" json:"delete_status"`
}

func (Role) TableName() string {
	return "mst_roles"
}
