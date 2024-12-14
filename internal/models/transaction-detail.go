package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	ContractNumber    uint    `gorm:"column:contract_number;unique;not null"`
	ConsumerID        uint    `gorm:"column:consumer_id"`
	CreditLimitID     uint    `gorm:"column:credit_limit_id"`
	OTR               float64 `gorm:"type:decimal(15,2);column:otr;comment:2 jenis: white godds(kulkas) dan total biaya adminstrasi dll(mobil)"`
	AdminFee          float64 `gorm:"type:decimal(15,2);column:admin_fee"`
	InstallmentAmount float64 `gorm:"type:decimal(15,2);column:installment_amount;comment:total cicilan perbulan"`
	InterestAmount    float64 `gorm:"type:decimal(15,2);column:interest_amount;comment:biaya++ (bunga)"`
	AssetName         string  `gorm:"type:decimal(15,2);column:asset_name"`
	PatnerID          uint    `gorm:"type:decimal(15,2);column:usage_source;comment:example:nama patner credit"`
}
