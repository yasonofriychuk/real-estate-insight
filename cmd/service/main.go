package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	serviceapi "github.com/yasonofriychuk/real-estate-insight/internal/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/html/index_page_handler"
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

	pg, err := postgres.New("postgres://postgres:password@postgres:5432/postgres")
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
	}
	defer pg.Close()

	osmStorage := osm_storage.New(pg.Pool)

	srv := serviceapi.API{
		BuildRoutesByPointsHandler:              build_routes_by_points.New(log, rb),
		ObjectsFindNearestInfrastructureHandler: objects_find_nearest_infrastructure.New(log, osmStorage, rb),
		IndexPageHandler:                        index_page_handler.New(log),
	}

	server, err := api.NewServer(srv)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create server")
		os.Exit(1)
	}

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to start server")
		os.Exit(1)
	}
}
