package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(100);not null;column:name;comment:mobil,rumah"`                             // Nama aset
	Type       string  `gorm:"type:varchar(100);not null;column:type;comment:Properti,Elektronik,Kendaraan,Investasi"` // Jenis aset
	Amount     float64 `gorm:"type:decimal(20,2);not null;column:amount"`                                              // Nilai total aset
	PartnerID  uint    `gorm:"column:partner_id;foreignkey:PartnerID;comment:partner could be null ,add null when this is from customer or conventional partner"`
	CustomerID uint    `gorm:"column:customer_id;foreignkey:CustomerID"`
}
