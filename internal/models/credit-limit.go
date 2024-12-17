package models

import (
	"gorm.io/gorm"
)

type CreditLimit struct {
	gorm.Model
	TenorMonths float64 `gorm:"column:tenor_months;comment:lama bulan pelunasan;not null"`
	LimitAmount float64 `gorm:"type:decimal(15,2);column:limit_amount;comment:limit transaction;not null"`
}

//credit limit explain
// Tenor dalam konteks ini mengacu pada jangka waktu pelunasan pinjaman atau kredit yang diberikan kepada konsumen untuk membayar suatu transaksi. Jadi, tenor 1, 2, 3, dan 6 menunjukkan pilihan durasi pembayaran dalam bulan. Penjelasannya sebagai berikut:

// Apa itu tenor?
// Tenor adalah istilah yang digunakan dalam dunia keuangan untuk menggambarkan periode waktu pembayaran cicilan hingga lunas.
// Dalam kasus ini:
// Tenor 1 bulan berarti konsumen harus melunasi pembayaran dalam waktu 1 bulan.
// Tenor 2 bulan berarti cicilan akan dilunasi dalam 2 bulan, dan seterusnya.
// Tenor 6 bulan menunjukkan pelunasan dalam waktu lebih panjang, yaitu 6 bulan.
// Mengapa tenor dimulai dari 1, 2, 3, dan 6 bulan?
// Flexibilitas: Rentang tenor ini memberi konsumen opsi untuk memilih durasi pembayaran sesuai kemampuan finansial mereka.
// Praktis & umum digunakan: Dalam dunia pinjaman atau pembiayaan, tenor pendek seperti 1-6 bulan sering digunakan untuk pembelian dengan nominal tidak terlalu besar.
// Tenor yang tidak tersedia, misalnya 4 atau 5 bulan, mungkin dihilangkan karena jarang diminati atau kebijakan perusahaan.
