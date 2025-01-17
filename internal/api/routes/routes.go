package routes

import (
	"authService/internal/api/controllers/authController"
	users2 "authService/internal/api/controllers/usersController"
	"authService/internal/api/routes/auth"
	"authService/internal/api/routes/users"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, usersController *users2.UsersController, authController *authController.AuthController) {
	users.RegisterUserRoutes(r, usersController)
	auth.RegisterAuthRoutes(r, authController)
}
