package repositories

import (
	"GO-PTTK/config"
	"GO-PTTK/models"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Create(project *models.Project) error
	GetList() ([]models.Project, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository() ProjectRepository {
	return &projectRepository{
		db: config.GetDB(),
	}
}

func (r *projectRepository) Create(project *models.Project) error {
	return r.db.Create(project).Error
}

func (r *projectRepository) GetList() ([]models.Project, error) {
	var projects []models.Project

	err := r.db.
		Where("status = ?", models.StatusDraft).
		Preload("Members").
		Preload("Attachments").
		Order("created_at desc").
		Find(&projects).Error

	return projects, err
}
