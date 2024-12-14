package models

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name       string  `gorm:"type:varchar(100);not null;column:name;comment:mobil,rumah"`                             // Nama aset
	Type       string  `gorm:"type:varchar(100);not null;column:type;comment:Properti,Elektronik,Kendaraan,Investasi"` // Jenis aset
	Amount     float64 `gorm:"type:decimal(20,2);not null;column:amount"`                                                // Nilai total aset
	PartnerID  uint    `gorm:"column:partner_id;comment:partner could be null"`
	CustomerID uint    `gorm:"column:customer_id;comment:only add this when this is from customer or conventional partner"`
}

// from: [official partner/conventional partner ]
// customerId:
