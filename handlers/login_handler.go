package handlers

import (
	"GO-PTTK/middlewares"
	"GO-PTTK/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginHandler struct {
	repo repositories.AdminLogin
}

func NewLoginHandler(repo repositories.AdminLogin) *loginHandler {
	return &loginHandler{
		repo: repo,
	}
}

func (h *loginHandler) Login(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid information!",
		})
		return
	}

	admin, err := h.repo.AuthenticateAdmin(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Incorrect username or password!",
		})
		return
	}

	token, err := middlewares.GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cannot login!!!",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"token",
		token,
		3600/60*15,
		"/",
		"",
		false, // secure = false vì đang dev (true nếu chạy https)
		true,  // http
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!!!",
		"admin":   admin,
		"token":   token,
	})
}
