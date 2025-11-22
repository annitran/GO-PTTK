package routes

import "github.com/gin-contrib/cors"

import (
	"GO-PTTK/handlers"
	"GO-PTTK/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Cho phép mở file upload
	router.Static("/static", "./uploads")

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
		},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	projectHandler := handlers.NewProjectHandler(repositories.NewProjectRepository())

	// USER SUBMIT PROJECT (không login)
	router.POST("/api/v1/projects/submit", projectHandler.SubmitProject)

	// STAFF (login)
	router.GET("/api/v1/projects/admin", projectHandler.AdminGetProjects)

	return router
}
