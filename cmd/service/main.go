package main

import (
	"context"
	"fmt"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/infrastructure_heatmap"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/location_list"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/profile_login"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_create"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_delete"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_edit"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_favorite"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_list"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/location"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/profile"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
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
	cfg := config.MustNewConfigWithEnv()

	jwtService := auth.NewJwtService(cfg.JwtSecretKey())
	rb := route_builder.NewRouteBuilder()

	pg, err := postgres.New(cfg.PgUrl())
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
	}
	defer pg.Close()

	developmentStorage := development.New(pg.Pool)
	infrastructureStorage := infrastructure.New(pg.Pool)
	coordinatesStorage := coordinates.New(pg.Pool)
	profileStorage := profile.New(pg.Pool)
	selectionStorage := selection.New(pg.Pool)
	locationStorage := location.New(pg.Pool)

	srv := serviceapi.API{
		BuildRoutesByPointsHandler:       build_routes_by_points.New(log, rb, coordinatesStorage),
		DevelopmentSearchHandler:         development_search_filter.New(log, developmentStorage),
		InfrastructureRadiusBoardHandler: infrastructure_radius_board.New(log, infrastructureStorage),
		HeatmapHandler:                   infrastructure_heatmap.New(log, infrastructureStorage),
		ProfileLoginHandler:              profile_login.New(log, jwtService, profileStorage),
		SelectionDeleteHandler:           selection_delete.New(log, selectionStorage),
		SelectionFavoriteHandler:         selection_favorite.New(log),
		SelectionCreateHandler:           selection_create.New(log, selectionStorage),
		SelectionListHandler:             selection_list.New(log, selectionStorage),
		LocationListHandler:              location_list.New(log, locationStorage),
		SelectionEditHandler:             selection_edit.New(log, selectionStorage),
	}

	server, err := api.NewServer(srv)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to create server")
		os.Exit(1)
	}

	corsAuthHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	}).Handler(auth.MustNewMiddleware(server, jwtService, api.UserLoginOperation))

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", corsAuthHandler))
	mux.Handle("/", http.FileServer(http.Dir("static")))

	log.WithContext(ctx).Info("server start")

	if err := http.ListenAndServe(
		fmt.Sprintf(":%s", cfg.HttpPort()),
		cors.AllowAll().Handler(mux),
	); err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to start server")
		os.Exit(1)
	}
}
