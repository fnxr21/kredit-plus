package models

import "gorm.io/gorm"

type Partner struct {
	gorm.Model
	Name              string `gorm:"type:varchar(50);column:name;not null"`
	Email             string `gorm:"type:varchar(50);column:email;not null"`
	PhoneNumber       string `gorm:"type:varchar(20);column:phone_number;not null"`
	BankAccount       string `gorm:"type:varchar(20);column:bank_account;not null"`
	AccountHolderName string `gorm:"type:varchar(50);column:account_holder_name;not null"`
	BankName          string `gorm:"type:varchar(20);column:bank_name;not null"`
	Status            int    `gorm:"type:tinyint(1);column:status;not null;default:'1';comment:0 =inactive,1 = active"`
}
