package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type Partner interface {
	CreatePartner(user models.Partner) (models.Partner, error)
	ListPartner() ([]models.Partner, error)
	PartnerByID(id uint) (models.Partner, error)
}

func RepositoryPartner(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePartner(user models.Partner) (models.Partner, error) {
	err := r.db.Create(&user).Scan(&user).Preload("PartnerBank").
		Error

	return user, err
}
func (r *repository) ListPartner() ([]models.Partner, error) {
	var partner []models.Partner
	err := r.db.Find(&partner).Preload("PartnerBank").
		Error

	return partner, err
}
func (r *repository) PartnerByID(id uint) (models.Partner, error) {
	var partner models.Partner
	err := r.db.First(&partner, id).Preload("PartnerBank").
		Error

	return partner, err
}
