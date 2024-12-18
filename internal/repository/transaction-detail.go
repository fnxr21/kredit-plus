package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type TransactionDetail interface {
	CreateTransactionDetail(user models.TransactionDetail) (models.TransactionDetail, error)
	ListTransactionDetail() ([]models.TransactionDetail, error)
	TransactionDetailByID(id uint) (models.TransactionDetail, error)
}

func RepositoryTransactionDetail(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateTransactionDetail(user models.TransactionDetail) (models.TransactionDetail, error) {
	err := r.db.Create(&user).Scan(&user).
		Error

	return user, err
}
func (r *repository) ListTransactionDetail() ([]models.TransactionDetail, error) {
	var detail []models.TransactionDetail
	err := r.db.Find(&detail).
		Error

	return detail, err
}
func (r *repository) TransactionDetailByID(id uint) (models.TransactionDetail, error) {
	var detail models.TransactionDetail
	err := r.db.First(&detail, "id=?", id).
		Error

	return detail, err
}
