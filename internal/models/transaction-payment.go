package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionPayment struct {
	gorm.Model
	TransactionDetailID uint              `gorm:"column:transaction_detail_id;foreignkey:TransactionID"`
	TransactionDetail   TransactionDetail `gorm:"foreignkey:TransactionDetailID"`
	PaymentDate         time.Time         `gorm:"column:payment_date;not null"` // Tanggal pembayaran
	Amount              float64           `gorm:"type:decimal(15,2);column:amount"`
	Status              string            `gorm:"type:enum('pending','completed','failed');column:status;default:'pending';not null"`
	PartnerID           uint              `gorm:"column:partner_id;comment:partner could be null ,add null when this is from customer or conventional partner"`
	Partner             Partner           `gorm:"foreignkey:PartnerID"`
	CustomerID          uint              `gorm:"column:customer_id"`
	Customer            Customer          `gorm:"foreignkey:CustomerID"`
}
