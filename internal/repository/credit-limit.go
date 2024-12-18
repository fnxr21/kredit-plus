package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type CreditLimit interface {
	CreateCreditLimit(user models.CreditLimit) (models.CreditLimit, error)
	ListCreditLimit() ([]models.CreditLimit, error)
	CreditLimitByID(id uint) (models.CreditLimit, error)
}

func RepositoryCreditLimit(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCreditLimit(user models.CreditLimit) (models.CreditLimit, error) {
	err := r.db.Create(&user).Scan(&user).
		Error

	return user, err
}
func (r *repository) ListCreditLimit() ([]models.CreditLimit, error) {
	var limit []models.CreditLimit
	err := r.db.Find(&limit).
		Error

	return limit, err
}
func (r *repository) CreditLimitByID(id uint) (models.CreditLimit, error) {
	var limit models.CreditLimit
	err := r.db.First(&limit).
		Error

	return limit, err
}
