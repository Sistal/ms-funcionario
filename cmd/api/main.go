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

	_ "github.com/Sistal/ms-funcionario/docs" // Swagger docs
)

// @title API Microservicio de Funcionarios
// @version 1.0
// @description API RESTful para la gestión integral de funcionarios en el sistema de control de uniformes empresariales
// @termsOfService http://swagger.io/terms/

// @contact.name Soporte API
// @contact.email soporte@sistal.cl

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @tag.name funcionarios
// @tag.description Operaciones CRUD y gestión de funcionarios

// @tag.name medidas
// @tag.description Gestión de medidas corporales de funcionarios

// @tag.name employees
// @tag.description Endpoints del BFF para funcionarios (perfil y medidas)

// @tag.name catálogos
// @tag.description Catálogos de datos maestros (cargos, géneros)

// @tag.name sucursales
// @tag.description Gestión de sucursales

// @tag.name traslados
// @tag.description Solicitudes de traslado

// @tag.name health
// @tag.description Health check del servicio

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Autenticación mediante token JWT. Formato: "Bearer {token}"

func main() {
	log.Println("Starting ms-funcionario application...")

	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Inicializar repositorios
	funcionarioRepo := repository.NewFuncionarioRepository(db)
	medidasRepo := repository.NewMedidasRepository(db)
	sucursalRepo := repository.NewSucursalRepository(db)
	cargoRepo := repository.NewCargoRepository(db)
	generoRepo := repository.NewGeneroRepository(db)

	// Inicializar servicios
	funcionarioService := services.NewFuncionarioService(funcionarioRepo, medidasRepo, sucursalRepo)
	catalogoService := services.NewCatalogoService(cargoRepo, generoRepo)

	// Inicializar handlers
	funcionarioHandler := handler.NewFuncionarioHandler(funcionarioService)
	profileHandler := handler.NewProfileHandler(funcionarioService)
	catalogoHandler := handler.NewCatalogoHandler(catalogoService)

	// Configurar router
	r := router.SetupRouter(funcionarioHandler, profileHandler, catalogoHandler, cfg.AllowedOrigins)

	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on port %s", cfg.ServerPort)

	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
