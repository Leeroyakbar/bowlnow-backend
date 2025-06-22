package models

import (
	"time"

	"github.com/google/uuid"
)

type Food struct {
	FoodID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"food_id"`
	CategoryID    uuid.UUID `gorm:"type:uuid" json:"category_id"`
	FoodName      string    `gorm:"type:varchar(100)" json:"food_name"`
	Description   string    `gorm:"type:text" json:"description"`
	Price         float64   `gorm:"type:numeric" json:"price"`
	Image         string    `gorm:"type:text" json:"image"`
	AvailableFlag int       `gorm:"type:int" json:"available_flag"` // 1 = tersedia, 0 = tidak
	DeleteStatus  string    `gorm:"type:varchar(10);default:'0'" json:"delete_status"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Category Category
}

// Untuk mapping ke tabel `mst_foods`
func (Food) TableName() string {
	return "mst_foods"
}
