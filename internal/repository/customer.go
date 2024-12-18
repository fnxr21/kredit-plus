package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type Customer interface {
	ListCustomer() ([]models.Customer, error)
	CustomerByID(id int) (models.Customer, error)
}

func RepositoryCustomer(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ListCustomer() ([]models.Customer, error) {
	var customer []models.Customer
	err := r.db.Find(&customer).
		Error

	return customer, err
}
func (r *repository) CustomerByID(id int) (models.Customer, error) {
	var customer models.Customer
	err := r.db.First(&customer).
		Error

	return customer, err
}
