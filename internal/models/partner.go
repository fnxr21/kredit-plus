package models

import "gorm.io/gorm"

type Partner struct {
	gorm.Model
	Name          string `gorm:"type:varchar(50);column:name;not null"`
	Email         string `gorm:"type:varchar(50);column:email;not null;unique"`
	PhoneNumber   string `gorm:"type:varchar(50);column:phone_number;not null"`
	Address       string `gorm:"type:varchar(50);column:address;not null"`
	PartnerBankID uint   `gorm:"column:partner_bank_id;foreignkey:PartnerBankID;not null"`
}
