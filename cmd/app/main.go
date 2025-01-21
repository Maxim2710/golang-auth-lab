package main

import (
	"github.com/Maxim2710/golang-auth-lab/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Application terminated with error: %v", err)
	}
}
