package auth

import (
	"authService/internal/api/controllers/authController"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authController *authController.AuthController) {

	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

}
