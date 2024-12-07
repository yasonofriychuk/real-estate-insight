package main

import (
	"context"
	"fmt"
	"github.com/yasonofriychuk/real-estate-insight/internal/config"
	"log/slog"
	"net/http"
	"os"

	serviceapi "github.com/yasonofriychuk/real-estate-insight/internal/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/objects/objects_find_nearest_infrastructure"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/routes/build_routes_by_points"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	osm_storage "github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/osm"
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

	osmStorage := osm_storage.New(pg.Pool)

	srv := serviceapi.API{
		BuildRoutesByPointsHandler:              build_routes_by_points.New(log, rb),
		ObjectsFindNearestInfrastructureHandler: objects_find_nearest_infrastructure.New(log, osmStorage, rb),
	}

	server, err := api.NewServer(srv)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create server")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", server))
	mux.Handle("/", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.HttpPort()), mux); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to start server")
		os.Exit(1)
	}
}
