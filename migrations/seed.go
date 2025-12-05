package migrations

import (
	"GO-PTTK/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

func SeedAdmin(db *gorm.DB) {
	var admin models.Admin

	// Kiểm tra admin có tồn tại chưa
	result := db.First(&admin, "username = ?", "admin")

	if result.Error == gorm.ErrRecordNotFound {
		// Tạo password hash
		hashed, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)

		newAdmin := models.Admin{
			Username: "admin",
			Password: string(hashed),
		}

		db.Create(&newAdmin)
		log.Println("Admin default created: admin / 123456")
	} else {
		log.Println("Admin already exists — skip")
	}
}
