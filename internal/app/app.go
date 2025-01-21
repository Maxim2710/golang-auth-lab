package app

import (
	"github.com/Maxim2710/golang-auth-lab/internal/config"
	"github.com/Maxim2710/golang-auth-lab/internal/database"
	"github.com/Maxim2710/golang-auth-lab/internal/database/repository"
	"github.com/Maxim2710/golang-auth-lab/internal/service"
	"github.com/Maxim2710/golang-auth-lab/internal/transport/http"
)

func Run() error {
	cfg, err := config.LoadConfig()

	if err != nil {
		return err
	}

	db, err := database.Connect(cfg.Database)

	if err != nil {
		return err
	}

	defer db.Close()

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)

	router := http.SetupRouter(authService)

	return router.Run("localhost:8080")
}
