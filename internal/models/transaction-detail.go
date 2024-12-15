package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	ContractNumber    uint    `gorm:"column:contract_number;unique;not null;unique"`
	OTR               float64 `gorm:"type:decimal(15,2);column:otr;comment:2 jenis: white godds(kulkas) dan total biaya adminstrasi dll(mobil)"`
	AdminFee          float64 `gorm:"type:decimal(15,2);column:admin_fee"`
	InstallmentAmount float64 `gorm:"type:decimal(15,2);column:installment_amount;comment:total cicilan perbulan"`
	InterestAmount    float64 `gorm:"type:decimal(15,2);column:interest_amount;comment:biaya++ (bunga)"`
	Status            string  `gorm:"type:enum('pending','accept','reject');column:payment;default:'pending';not null"`
	CreditLimitID     uint    `gorm:"column:credit_limit_id;foreignkey:CreditLimitID"`
	PartnerBankID     uint    `gorm:"column:partner_bank_id;foreignkey:PartnerBankID;not null"`
	AssetID           uint    `gorm:"column:asset_id;foreignkey:AssetID"`
	PartnerID         uint    `gorm:"column:partner_id;foreignkey:PartnerID;comment:partner could be null ,add null when this is from customer or conventional partner"`
	CustomerID        uint    `gorm:"column:customer_id;foreignkey:CustomerID"`
}
