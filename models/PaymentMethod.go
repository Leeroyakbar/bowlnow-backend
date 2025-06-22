package models

import "github.com/google/uuid"

type PaymentMethod struct {
	PaymentID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"payment_id"`
	PaymentName  string    `gorm:"type:varchar(30)" json:"payment_name"`
	DeleteStatus string    `gorm:"type:varchar(10);default:'0'" json:"delete_status"`
}
