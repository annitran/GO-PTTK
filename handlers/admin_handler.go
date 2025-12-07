package handlers

import (
	"GO-PTTK/models"
	"GO-PTTK/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type adminHandler struct {
	repo repositories.AdminRepository
}

func NewAdminHandler(repo repositories.AdminRepository) *adminHandler {
	return &adminHandler{repo: repo}
}

func GetAdmin(c *gin.Context) {
	adminData, exists := c.Get("admin")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Admin not found!",
		})
		return
	}

	admin := adminData.(*models.Admin)

	c.JSON(http.StatusOK, gin.H{
		"admin": admin,
	})
}
