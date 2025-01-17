package users

import (
	"authService/internal/api/controllers/usersController"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, usersController *usersController.UsersController) {
	usersRoutes := r.Group("/users")
	usersRoutes.Use(usersController.AuthMiddleware.JWTMiddleware())
	{
		usersRoutes.GET("/getUserById", usersController.GetUserByID)
		usersRoutes.GET("/getUserByEmail", usersController.GetUserByEmail)
		usersRoutes.GET("/getUserByUsername", usersController.GetUserByUsername)
		usersRoutes.POST("/createUser", usersController.CreateUser)
		usersRoutes.POST("/updateUser", usersController.UpdateUser)
		usersRoutes.POST("/deleteUserById", usersController.DeleteUserById)
	}

}
