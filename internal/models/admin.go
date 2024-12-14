package models

import "gorm.io/gorm"

type MyUser struct {
	gorm.Model
	Username    string `gorm:"type:varchar(50);column:username;not null"`
	Password    string `gorm:"type:varchar(50);column:password;not null"`
	PhoneNumber string `gorm:"type:varchar(20);column:phone_number;not null"`
	Email       string `gorm:"type:varchar(50);column:email;not null"`
}
