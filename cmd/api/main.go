package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/xoesae/cid-api/internal/config"
	"github.com/xoesae/cid-api/internal/domain/repository"
	"github.com/xoesae/cid-api/internal/domain/service"
	"github.com/xoesae/cid-api/internal/infra/database/pg"
	"github.com/xoesae/cid-api/internal/infra/http"
	"github.com/xoesae/cid-api/internal/infra/http/handler"
	"github.com/xoesae/cid-api/pkg/logger"
	"log/slog"
	"os"
)

func main() {
	// config
	cfg := config.GetConfig()

	// setup logger
	logger.Init(cfg.LogLevel)

	// router
	r := chi.NewRouter()

	// database connection
	conn, err := pg.NewConnection(cfg)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		panic(err)
	}
	defer conn.Close()

	// repositories
	categoryRepo := repository.NewCategoryRepository(conn.DB())
	subcategoryRepo := repository.NewSubcategoryRepository(conn.DB())

	// services
	cidService := service.NewCidService(categoryRepo, subcategoryRepo)

	// handlers
	cidHandler := handler.CidHandler{
		CidService: cidService,
	}

	// register handler
	http.RegisterCidRoutes(r, &cidHandler)

	_server := http.New(http.Config{
		Port:   cfg.Port,
		Router: r,
	})

	err = _server.Start()
	if err != nil {
		slog.Error("failed to start server", err)
		os.Exit(1)
	}
}
