package repositories

import (
	"kredit-plus/internal/models"

	"gorm.io/gorm"
)

type Asset interface {
	CreateAsset(user models.Asset) (models.Asset, error)
	ListAsset() ([]models.Asset, error)
	AssetByID(id uint) (models.Asset, error)
}

func RepositoryAsset(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAsset(user models.Asset) (models.Asset, error) {
	err := r.db.Create(&user).Scan(&user).
		Error

	return user, err
}
func (r *repository) ListAsset() ([]models.Asset, error) {
	var asset []models.Asset
	err := r.db.Find(&asset).
		Error

	return asset, err
}
func (r *repository) AssetByID(id uint) (models.Asset, error) {
	var asset models.Asset
	err := r.db.First(&asset).
		Error

	return asset, err
}
