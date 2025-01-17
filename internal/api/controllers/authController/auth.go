package authController

import (
	"authService/internal/api/controllers/authController/dto"
	"authService/internal/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *auth.Service
}

func NewController(authService *auth.Service) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (a *AuthController) Login(ctx *gin.Context) {
	var loginReq dto.LoginReq
	if err := ctx.ShouldBindBodyWithJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseChan := make(chan *dto.ReturnToken)

	go a.authService.Login(loginReq, responseChan)

	tokenObj := <-responseChan

	err := tokenObj.Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tokenObj.Token)

}

func (a *AuthController) Register(ctx *gin.Context) {
	var registerReq dto.RegisterReq
	if err := ctx.ShouldBindBodyWithJSON(&registerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseChan := make(chan *dto.ReturnConfirmation)

	go a.authService.Register(registerReq, responseChan)

	confirmObj := <-responseChan

	if !confirmObj.Confirmation {
		ctx.JSON(http.StatusOK, gin.H{"error": confirmObj.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})

}
