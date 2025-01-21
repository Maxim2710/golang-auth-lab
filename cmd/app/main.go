package main

import (
	"github.com/Maxim2710/golang-auth-lab/internal/app"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.DebugMode)

	if err := app.Run(); err != nil {
		log.Fatalf("Application terminated with error: %v", err)
	}
}
