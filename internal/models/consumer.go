package models

import (
	"time"

	"gorm.io/gorm"
)

type Consumer struct {
	gorm.Model
	Nik         string    `gorm:"type:varchar(50);column:nik;unique;not null"`
	FullName    string    `gorm:"type:varchar(50);column:full_name;not null"`
	LegalName   string    `gorm:"type:varchar(50);column:legal_name;not null"`
	Birthplace  string    `gorm:"type:varchar(50);column:legal_name;not null"`
	BirthDate   time.Time `gorm:"type:date;column:dob;not null"`
	Salary      float64   `gorm:"type:decimal(15,2);column:salary;not null"`
	ImageKTP    string    `gorm:"type:varchar(50);column:image_ktp;comment:pathfile;not null"`
	ImageSelfie string    `gorm:"type:varchar(50);column:image_selfie;comment:pathfile;not null"`
	PatnerID    uint      `gorm:"column:patner_id"`
}
