package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/rs/cors"

	serviceapi "github.com/yasonofriychuk/real-estate-insight/internal/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/build_routes_by_points"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/development_search_filter"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/infrastructure_radius_board"
	"github.com/yasonofriychuk/real-estate-insight/internal/config"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/coordinates"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/infrastructure"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/postgres"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)
	rb := route_builder.NewRouteBuilder()
	cfg := config.MustNewConfigWithEnv()

	pg, err := postgres.New(cfg.PgUrl())
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
	}
	defer pg.Close()

	developmentStorage := development.New(pg.Pool)
	infrastructureStorage := infrastructure.New(pg.Pool)
	coordinatesStorage := coordinates.New(pg.Pool)

	srv := serviceapi.API{
		BuildRoutesByPointsHandler:       build_routes_by_points.New(log, rb, coordinatesStorage),
		DevelopmentSearchHandler:         development_search_filter.New(log, developmentStorage),
		InfrastructureRadiusBoardHandler: infrastructure_radius_board.New(log, infrastructureStorage),
	}

	server, err := api.NewServer(srv)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create server")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", server))
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.WithContext(ctx).Info("server start")

	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", cfg.HttpPort()),
		cors.New(cors.Options{AllowedOrigins: []string{}}).Handler(mux),
	); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to start server")
		os.Exit(1)
	}
}
