package repositories

import (
	"GO-PTTK/config"
	"GO-PTTK/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*models.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{
		db: config.GetDB(),
	}
}

func (r *adminRepository) FindByUsername(username string) (*models.Admin, error) {
	var admin models.Admin

	if err := r.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}
