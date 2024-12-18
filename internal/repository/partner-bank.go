package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type PartnerBank interface {
	CreatePartnerBank(user models.PartnerBank) (models.PartnerBank, error)
	ListPartnerBank() ([]models.PartnerBank, error)
	PartnerBankByID(id uint) (models.PartnerBank, error)
}

func RepositoryPartnerBank(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePartnerBank(user models.PartnerBank) (models.PartnerBank, error) {
	err := r.db.Create(&user).Scan(&user).
		Error

	return user, err
}
func (r *repository) ListPartnerBank() ([]models.PartnerBank, error) {
	var bank []models.PartnerBank
	err := r.db.Find(&bank).
		Error

	return bank, err
}
func (r *repository) PartnerBankByID(id uint) (models.PartnerBank, error) {
	var bank models.PartnerBank
	err := r.db.First(&bank).
		Error

	return bank, err
}
