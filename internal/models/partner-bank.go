package models

import "gorm.io/gorm"

type PartnerBank struct {
	gorm.Model
	BankAccount       string `gorm:"type:varchar(50);column:bank_account;not null;unique"`
	AccountHolderName string `gorm:"type:varchar(50);column:account_holder_name;not null"`
	BankName          string `gorm:"type:varchar(50);column:bank_name;not null"`
}
