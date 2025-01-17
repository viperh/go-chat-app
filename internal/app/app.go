package app

import (
	"authService/internal/api/controllers/authController"
	"authService/internal/api/controllers/usersController"
	"authService/internal/api/middlewares"
	"authService/internal/api/routes"
	"authService/internal/config"
	"authService/internal/pkg/auth"
	"authService/internal/pkg/token"
	users2 "authService/internal/pkg/users"
	"authService/internal/rlog"
	"authService/internal/storage/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
)

type App struct {
	cfg            *config.Config
	gin            *gin.Engine
	UserController *usersController.UsersController
	UserProvider   *users2.Service
	UserStorage    *postgres.Postgres
	AuthController *authController.AuthController
	AuthProvider   *auth.Service
}

func NewApp() *App {
	cfg := config.NewConfig("dev")
	g := gin.Default()

	usersStorage := postgres.NewDatabase(cfg)
	usersService := users2.NewService(usersStorage)
	tokenService := token.NewService(cfg)
	authMiddleware := middlewares.NewAuthMiddleware(cfg, tokenService)
	usersControllerVar := usersController.NewController(usersService, authMiddleware)
	authService := auth.NewService(usersStorage, tokenService)
	authControllerVar := authController.NewController(authService)
	routes.RegisterRoutes(g, usersControllerVar, authControllerVar)

	return &App{
		cfg:            cfg,
		gin:            g,
		UserController: usersControllerVar,
		UserProvider:   usersService,
		UserStorage:    usersStorage,
		AuthController: authControllerVar,
		AuthProvider:   authService,
	}

}

func (a *App) Run() {
	log := rlog.NewLogger("dev", "APP")

	log.Info("Server started successfully!")

	err := a.gin.Run(fmt.Sprintf(":%s", a.cfg.ServerPort))

	if err != nil {
		log.Fatal("Could not start server")
	}

}
