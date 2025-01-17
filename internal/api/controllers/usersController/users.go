package usersController

import (
	"authService/internal/api/controllers/usersController/dto"
	"authService/internal/api/middlewares"
	"authService/internal/models"
	"authService/internal/pkg/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UsersController struct {
	Service        *users.Service
	AuthMiddleware *middlewares.AuthMiddleware
}

func NewController(service *users.Service, authMiddleware *middlewares.AuthMiddleware) *UsersController {
	return &UsersController{
		Service:        service,
		AuthMiddleware: authMiddleware,
	}
}

func (s *UsersController) GetUserByID(c *gin.Context) {
	var id dto.ID
	if err := c.ShouldBindBodyWithJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseChan := make(chan *dto.ReturnUser)

	go s.Service.GetUserByID(id.ID, responseChan)

	userObj := <-responseChan

	err := userObj.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userObj.User)
}

func (s *UsersController) GetUserByEmail(c *gin.Context) {
	var email dto.Email
	if err := c.ShouldBindBodyWithJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseChan := make(chan *dto.ReturnUser)

	go s.Service.GetUserByEmail(email.Email, responseChan)

	userObj := <-responseChan

	err := userObj.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userObj.User)

}

func (s *UsersController) GetUserByUsername(c *gin.Context) {
	var username dto.Username
	if err := c.ShouldBindBodyWithJSON(&username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseChan := make(chan *dto.ReturnUser)

	go s.Service.GetUserByUsername(username.Username, responseChan)

	userObj := <-responseChan

	err := userObj.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userObj.User)
}

func (s *UsersController) CreateUser(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser := &models.User{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}

	responseChan := make(chan *dto.ReturnConfirmation)

	go s.Service.CreateUser(createdUser, responseChan)

	confirmObj := <-responseChan

	if !confirmObj.Confirmation {
		c.JSON(http.StatusInternalServerError, gin.H{"error": confirmObj.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})

}

func (s *UsersController) UpdateUser(c *gin.Context) {
	var user dto.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser := &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}

	responseChan := make(chan *dto.ReturnConfirmation)

	go s.Service.UpdateUser(updatedUser, responseChan)

	confirmObj := <-responseChan

	if !confirmObj.Confirmation {
		c.JSON(http.StatusInternalServerError, gin.H{"error": confirmObj.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})

}

func (s *UsersController) DeleteUserById(c *gin.Context) {
	var id dto.ID
	if err := c.ShouldBindBodyWithJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseChan := make(chan *dto.ReturnConfirmation)

	go s.Service.DeleteUser(id.ID, responseChan)

	confirmObj := <-responseChan

	if !confirmObj.Confirmation {
		c.JSON(http.StatusInternalServerError, gin.H{"error": confirmObj.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
