package main

import (
	"fmt"
	"log"

	"github.com/Sistal/ms-funcionario/config"
	"github.com/Sistal/ms-funcionario/internal/application/services"
	"github.com/Sistal/ms-funcionario/internal/infrastructure/database"
	"github.com/Sistal/ms-funcionario/internal/infrastructure/repository"
	"github.com/Sistal/ms-funcionario/internal/interfaces/http/handler"
	"github.com/Sistal/ms-funcionario/internal/interfaces/http/router"
)

func main() {
	log.Println("Starting ms-funcionario application...")

	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	funcionarioRepo := repository.NewFuncionarioRepository(db)
	funcionarioService := services.NewFuncionarioService(funcionarioRepo)
	funcionarioHandler := handler.NewFuncionarioHandler(funcionarioService)

	r := router.SetupRouter(funcionarioHandler)

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on port %s", cfg.ServerPort)

	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
