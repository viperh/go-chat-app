package main

import (
	app2 "authService/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := app2.NewApp()
	app.Run()
}
