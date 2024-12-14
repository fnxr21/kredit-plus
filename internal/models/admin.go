package models

import "gorm.io/gorm"

type MyUser struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);column:username"`
	Password string `gorm:"type:varchar(50);column:password"`
	Email    string `gorm:"type:varchar(50);column:email"`
}
