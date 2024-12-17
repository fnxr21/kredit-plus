package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type AdminAuth interface {
	Register(user models.MyUser) (models.MyUser, error)
	Login(username string) (models.MyUser, error)
	Reauth(id uint) (models.MyUser, error)
}

func RepositoryAdminAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.MyUser) (models.MyUser, error) {
	err := r.db.Create(&user).Scan(&user).
		Error

	return user, err
}
func (r *repository) Login(username string) (models.MyUser, error) {
	var user models.MyUser
	err := r.db.First(&user, "username=?", username).
		Error

	return user, err
}
func (r *repository) Reauth(id uint) (models.MyUser, error) {
	var user models.MyUser
	err := r.db.First(&user, id).Error
	return user, err
}
