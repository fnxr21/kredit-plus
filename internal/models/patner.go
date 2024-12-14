package models

import "gorm.io/gorm"

type Patner struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);column:name;not null"`
	Username  string `gorm:"type:varchar(50);column:username;not null"`
	Password  string `gorm:"type:varchar(50);column:password;not null"`
	Email     string `gorm:"type:varchar(50);column:email;not null"`
	Handphone string `gorm:"type:varchar(50);column:handphone;not null"`
}
