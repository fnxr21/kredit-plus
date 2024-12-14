package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionPayment struct {
	gorm.Model
	TransactionID uint      `gorm:"column:transaction_id"`
	PaymentDate   time.Time `gorm:"column:payment_date;not null"` // Tanggal pembayaran
	Amount        float64   `gorm:"type:decimal(15,2);column:amount"`
	Status        string    `gorm:"type:enum('pending','completed','failed');column:status;default:'pending';not null"`
}
