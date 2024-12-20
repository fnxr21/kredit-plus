package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type CustomerAuth interface {
	RegisterCustomer(user models.Customer) (models.Customer, error)
	LoginCustomer(username string) (models.Customer, error)
	ReauthCustomer(id uint) (models.Customer, error)
}

func RepositoryCustomerAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RegisterCustomer(user models.Customer) (models.Customer, error) {
	err := r.db.Create(&user).Scan(&user).
		Error

	return user, err
}

func (r *repository) LoginCustomer(username string) (models.Customer, error) {
	var user models.Customer
	err := r.db.First(&user, "username=?", username).
		Error

	return user, err
}
func (r *repository) ReauthCustomer(id uint) (models.Customer, error) {
	var user models.Customer
	err := r.db.First(&user, "id=?", id).Error
	return user, err
}
