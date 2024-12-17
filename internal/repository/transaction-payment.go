package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type TransactionPayment interface {
	CreateTransactionPayment(user models.TransactionPayment) (models.TransactionPayment, error)
	ListTransactionPayment() ([]models.TransactionPayment, error)
	TransactionPaymentByID(id int) (models.TransactionPayment, error)
}

func RepositoryTransactionPayment(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransactionPayment(user models.TransactionPayment) (models.TransactionPayment, error) {
	err := r.db.Create(&user).Scan(&user).
		Error

	return user, err
}
func (r *repository) ListTransactionPayment() ([]models.TransactionPayment, error) {
	var payment []models.TransactionPayment
	err := r.db.Find(&payment).
		Error

	return payment, err
}
func (r *repository) TransactionPaymentByID(id int) (models.TransactionPayment, error) {
	var payment models.TransactionPayment
	err := r.db.First(&payment).
		Error

	return payment, err
}
