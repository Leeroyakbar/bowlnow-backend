package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"transaction_id"`
	UserID          uuid.UUID `gorm:"type:uuid" json:"user_id"`
	DeliveryAddress string    `gorm:"type:text" json:"delivery_address"`
	PaymentMethodID uuid.UUID `gorm:"type:uuid" json:"payment_method_id"`
	TrxAmount       float64   `gorm:"type:numeric" json:"trx_amount"`
	FeeAmount       float64   `gorm:"type:numeric" json:"fee_amount"`
	GrandTotal      float64   `gorm:"type:numeric" json:"grand_total"`
	PhoneNumber     string    `gorm:"type:varchar(20)" json:"phone_number"`
	Status          string    `gorm:"type:varchar(20)" json:"status"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// relasi
	User          User
	PaymentMethod PaymentMethod
}

func (Transaction) TableName() string {
	return "trx_transactions"
}
