package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Username    string    `gorm:"type:varchar(100);column:username;not null;unique"`
	Password    string    `gorm:"type:varchar(100);column:password;not null"`
	Email       string    `gorm:"type:varchar(100);column:email;not null;unique"`
	PhoneNumber string    `gorm:"type:varchar(100);column:phone_number;not null"`
	Nik         string    `gorm:"type:varchar(100);column:nik;not null;unique"`
	FullName    string    `gorm:"type:varchar(100);column:full_name;not null"`
	LegalName   string    `gorm:"type:varchar(100);column:legal_name;not null"`
	Birthplace  string    `gorm:"type:varchar(100);column:birthplace;not null"`
	BirthDate   time.Time `gorm:"type:date;column:dob;not null"`
	Salary      float64   `gorm:"type:decimal(15,2);column:salary;not null"`
	ImageKTP    string    `gorm:"type:varchar(255);column:image_ktp;comment:pathfile;not null"`
	ImageSelfie string    `gorm:"type:varchar(255);column:image_selfie;comment:pathfile;not null"`
}
