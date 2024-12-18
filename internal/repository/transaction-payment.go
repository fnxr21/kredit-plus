package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type TransactionPayment interface {
	CreateTransactionPayment(user models.TransactionPayment) (models.TransactionPayment, error)
	ListTransactionPayment() ([]models.TransactionPayment, error)
	TransactionPaymentByID(id uint) (models.TransactionPayment, error)
}

func RepositoryTransactionPayment(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransactionPayment(user models.TransactionPayment) (models.TransactionPayment, error) {
	// Inisialisasi transaksi
	tx := r.db.Begin()
	if tx.Error != nil {
		return user, tx.Error
	}

	// Coba insert data
	err := tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return user, err
	}

	// Commit transaksi jika berhasil
	if err := tx.Commit().Error; err != nil {
		return user, err
	}
	return user, err
}
func (r *repository) ListTransactionPayment() ([]models.TransactionPayment, error) {
	var payment []models.TransactionPayment
	// implementation offsite
	// .Offset(100).Limit(limit)
	err := r.db.Find(&payment).
		Error

	return payment, err
}
func (r *repository) TransactionPaymentByID(id uint) (models.TransactionPayment, error) {
	var payment models.TransactionPayment
	err := r.db.First(&payment, "id=?", id).
		Error

	return payment, err
}
