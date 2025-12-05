package migrations

import (
	"GO-PTTK/config"
	"GO-PTTK/models"
	"fmt"
	"log"
)

func Migrate() {
	db := config.GetDB()

	err := db.AutoMigrate(
		&models.Project{},
		&models.ProjectMember{},
		&models.ProjectAttachment{},
		&models.ProjectReview{},
		&models.Admin{},
	)

	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	fmt.Println("Database migrated successfully!")
}
