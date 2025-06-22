package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	CategoryID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"category_id"`
	CategoryName string    `gorm:"type:varchar(100)" json:"category_name"`
	DeleteStatus string    `gorm:"type:varchar(10);default:'0'" json:"delete_status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// Mapping ke tabel `mst_categories`
func (Category) TableName() string {
	return "mst_categories"
}
