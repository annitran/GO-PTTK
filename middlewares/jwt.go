package middlewares

import (
	"GO-PTTK/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func AuthToken(repo repositories.AdminRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy token từ cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Not logged in!",
			})
			c.Abort()
			return
		}

		// Parse và kiểm tra token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired token!",
			})
			c.Abort()
			return
		}

		admin, err := repo.FindByUsername(claims.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Admin not found!",
			})
			c.Abort()
			return
		}

		// Lưu thông tin admin vào context
		c.Set("admin", admin)

		c.Next()
	}
}
