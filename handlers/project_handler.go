package handlers

import (
	"GO-PTTK/models"
	"GO-PTTK/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type projectHandler struct {
	projectRepo repositories.ProjectRepository
}

func NewProjectHandler(projectRepo repositories.ProjectRepository) *projectHandler {
	return &projectHandler{projectRepo: projectRepo}
}

// User submit project (không cần đăng nhập)
func (h *projectHandler) SubmitProject(c *gin.Context) {
	// Lấy thông tin cơ bản
	title := c.PostForm("title")
	proposer := c.PostForm("proposer_name")
	email := c.PostForm("email")
	field := c.PostForm("field")

	if title == "" || proposer == "" || email == "" || field == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Parse ngày bắt đầu - kết thúc
	var startPtr, endPtr *time.Time
	if t := c.PostForm("expected_start"); t != "" {
		tt, _ := time.Parse("2006-01-02", t)
		startPtr = &tt
	}
	if t := c.PostForm("expected_end"); t != "" {
		tt, _ := time.Parse("2006-01-02", t)
		endPtr = &tt
	}

	// Tạo struct đề tài
	project := models.Project{
		Title:         title,
		ProposerName:  proposer,
		Email:         email,
		Field:         field,
		ExpectedStart: startPtr,
		ExpectedEnd:   endPtr,
		Status:        models.StatusDraft,
	}

	// Thành viên
	memberNames := c.PostFormArray("members[]")
	memberRoles := c.PostFormArray("roles[]")

	for i := range memberNames {
		project.Members = append(project.Members, models.ProjectMember{
			Name: memberNames[i],
			Role: memberRoles[i],
		})
	}

	// Upload file
	form, _ := c.MultipartForm()
	files := form.File["files"]

	for _, file := range files {
		path := "uploads/" + file.Filename
		c.SaveUploadedFile(file, path)

		project.Attachments = append(project.Attachments, models.ProjectAttachment{
			FileName: file.Filename,
			FileURL:  "/static/" + file.Filename,
		})
	}

	// Lưu DB
	if err := h.projectRepo.Create(&project); err != nil {
		c.JSON(500, gin.H{"error": "Cannot save project"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Project submitted successfully",
		"project": project,
	})
}

// Admin
func (h *projectHandler) AdminGetProjects(c *gin.Context) {
	projects, err := h.projectRepo.GetList()
	if err != nil {
		c.JSON(500, gin.H{"error": "Cannot load project list"})
		return
	}

	c.JSON(200, gin.H{"projects": projects})
}
