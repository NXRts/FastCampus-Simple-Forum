package middleware

import (
	"errors"
	"net/http"
	"strings"

	config "github.com/NXRts/fsatcampus/internal/configs"
	"github.com/NXRts/fsatcampus/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := config.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("userID", userId)
		c.Set("username", username)
		c.Next()
	}
}
