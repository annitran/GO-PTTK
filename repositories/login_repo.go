package repositories

import (
	"GO-PTTK/config"
	"GO-PTTK/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminLogin interface {
	AuthenticateAdmin(username, password string) (*models.Admin, error)
}

type adminLogin struct {
	db *gorm.DB
}

func NewAdminLogin() AdminLogin {
	return &adminLogin{
		db: config.GetDB(),
	}
}

func (r *adminLogin) AuthenticateAdmin(username, password string) (*models.Admin, error) {
	var admin models.Admin

	if err := r.db.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &admin, nil
}
