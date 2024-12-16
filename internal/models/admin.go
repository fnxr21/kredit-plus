package models

import "gorm.io/gorm"

type MyUser struct {
	gorm.Model
	Username    string `gorm:"type:varchar(100);column:username;not null;unique"`
	Password    string `gorm:"type:varchar(100);column:password;not null"`
	PhoneNumber string `gorm:"type:varchar(100);column:phone_number;not null"`
	Email       string `gorm:"type:varchar(100);column:email;not null;unique"`
}
