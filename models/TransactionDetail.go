package models

import (
	"time"

	"github.com/google/uuid"
)

type TransactionDetail struct {
	TrxDetailID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"trx_detail_id"`
	TransactionID uuid.UUID `gorm:"type:uuid" json:"transaction_id"`
	FoodID        uuid.UUID `gorm:"type:uuid" json:"food_id"`
	Quantity      int       `gorm:"type:int" json:"quantity"`
	Price         float64   `gorm:"type:numeric" json:"price"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Transaction Transaction
	Food        Food
}

func (TransactionDetail) TableName() string {
	return "trx_transaction_details"
}
