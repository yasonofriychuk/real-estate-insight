package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	serviceapi "github.com/yasonofriychuk/real-estate-insight/internal/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/routes/build_routes_by_points"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/postgres"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)
	rb := route_builder.NewRouteBuilder()

	pg, err := postgres.New("postgres://osmuser:osmpassword@localhost:5432/osm")
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
	}
	defer pg.Close()

	buildRoutesByPointsHandler := build_routes_by_points.New(log, rb)

	srv := serviceapi.API{
		BuildRoutesByPointsHandler: buildRoutesByPointsHandler,
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
