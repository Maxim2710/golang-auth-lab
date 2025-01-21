package app

import (
	"github.com/Maxim2710/golang-auth-lab/internal/config"
	"github.com/Maxim2710/golang-auth-lab/internal/database"
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

	return nil
}
