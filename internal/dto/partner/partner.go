package partnerdto

import "gorm.io/gorm"

type Partner struct {
	gorm.Model
	Name          string `gorm:"type:varchar(100);column:name;not null"`
	Email         string `gorm:"type:varchar(100);column:email;not null;unique"`
	PhoneNumber   string `gorm:"type:varchar(100);column:phone_number;not null"`
	Address       string `gorm:"type:varchar(100);column:address;not null"`
	PartnerBankID uint   `gorm:"column:partner_bank_id;not null"`
}
