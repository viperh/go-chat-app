package middlewares

import (
	"authService/internal/config"
	"authService/internal/pkg/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	secretKey    string
	tokenService *token.Service
}

func NewAuthMiddleware(cfg *config.Config, tokenService *token.Service) *AuthMiddleware {
	return &AuthMiddleware{
		secretKey:    cfg.JwtKey,
		tokenService: tokenService,
	}
}

func (a *AuthMiddleware) JWTMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := a.tokenService.CheckToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()

	}

}
