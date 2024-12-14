package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	ContractNumber    uint    `gorm:"column:contract_number;unique;not null"`
	OTR               float64 `gorm:"type:decimal(15,2);column:otr;comment:2 jenis: white godds(kulkas) dan total biaya adminstrasi dll(mobil)"`
	AdminFee          float64 `gorm:"type:decimal(15,2);column:admin_fee"`
	InstallmentAmount float64 `gorm:"type:decimal(15,2);column:installment_amount;comment:total cicilan perbulan"`
	InterestAmount    float64 `gorm:"type:decimal(15,2);column:interest_amount;comment:biaya++ (bunga)"`
	CreditLimitID     uint    `gorm:"column:credit_limit_id"`
	AssetID           uint    `gorm:"column:asset_id"`
	ConsumerID        uint    `gorm:"column:consumer_id"`
	PatnerID          uint    `gorm:"column:partner_id"`
}
