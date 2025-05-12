package main

import (
	"fmt"
	"github.com/xoesae/cid-api/internal/config"
	"github.com/xoesae/cid-api/internal/domain/repository"
	"github.com/xoesae/cid-api/internal/domain/service"
	"github.com/xoesae/cid-api/internal/infra/database/pg"
	"log/slog"
)

func main() {
	// config
	cfg := config.GetConfig()

	// access the host out of the container
	cfg.DbHost = cfg.DbHostCli

	// database connection
	conn, err := pg.NewConnection(cfg)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		panic(err)
	}
	defer conn.Close()

	// repositories
	chapterRepo := repository.NewChapterRepository(conn.DB())
	groupRepo := repository.NewGroupRepository(conn.DB())
	categoryRepo := repository.NewCategoryRepository(conn.DB())
	subcategoryRepo := repository.NewSubcategoryRepository(conn.DB())

	// services
	importer := service.NewImportService(chapterRepo, groupRepo, categoryRepo, subcategoryRepo)

	xmlPath := "resources/CID10.xml"
	if err := importer.RunImport(xmlPath); err != nil {
		panic(fmt.Sprintf("import failed: %v", err))
	}
}
