package models

import "github.com/google/uuid"

// ✅ models/role.go
type Role struct {
	RoleID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"role_id"`
	RoleName     string    `gorm:"type:varchar(50);unique" json:"role_name"`
	DeleteStatus int       `gorm:"type:smallint" json:"delete_status"`
}

func (Role) TableName() string {
	return "mst_roles"
}
