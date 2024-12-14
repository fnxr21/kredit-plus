package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionPayment struct {
	gorm.Model
	TransactionID uint
	PaymentDate   time.Time `gorm:"not null"` // Tanggal pembayaran
	Amount        float64
	PaymentStatus string `gorm:"type:enum('pending','completed','failed');default:'pending';not null"`
}
