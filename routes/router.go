package routes

import (
	"GO-PTTK/handlers"
	"GO-PTTK/middlewares"
	"GO-PTTK/repositories"

	"github.com/gin-contrib/cors"
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

	// User
	projectHandler := handlers.NewProjectHandler(repositories.NewProjectRepository())
	router.POST("/api/v1/projects/submit", projectHandler.SubmitProject)

	// Login
	loginHandler := handlers.NewLoginHandler(repositories.NewAdminLogin())
	router.POST("/api/v1/login", loginHandler.Login)

	// Admin
	adminRepo := repositories.NewAdminRepository()
	auth := router.Group("/api/v1/admin", middlewares.AuthToken(adminRepo))
	{
		auth.GET("/projects", projectHandler.AdminGetProjects)
	}

	return router
}
