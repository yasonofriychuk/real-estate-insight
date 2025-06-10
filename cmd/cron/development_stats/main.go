package main

import (
	"context"
	"github.com/yasonofriychuk/real-estate-insight/internal/config"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/infrastructure"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/postgres"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
	"log/slog"
	"os"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)
	cfg := config.MustNewConfigWithEnv()

	pg, err := postgres.New(cfg.PgUrl())
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to connect to postgres")
		os.Exit(1)
	}
	defer pg.Close()

	developmentStorage := development.New(pg.Pool)
	infrastructureStorage := infrastructure.New(pg.Pool)

	rb := route_builder.NewRouteBuilder()

	devs, err := developmentStorage.SearchDevelopmentByFilters(ctx, development.Filter{}, nil)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to search development by filters")
		os.Exit(1)
	}

	for _, dev := range devs {
		nearestInfrastructure, err := infrastructureStorage.NearestInfrastructure(ctx, dev.Coordinates.Point, 5)
		if err != nil {
			log.WithContext(ctx).WithError(err).Error("failed to nearest infrastructure")
			continue
		}

		nearestInfrastructureStats := make(map[osm.ObjType]float64, 6)
		for infraType, infraObjs := range nearestInfrastructure {
			for _, obj := range infraObjs {
				routes, err := rb.BuildRoute(dev.Coordinates.Point, obj.Coordinates.Point, route_builder.FootTransportType)
				if err != nil {
					log.WithContext(ctx).WithError(err).WithFields(map[string]any{
						"from":  dev.Coordinates.Point,
						"to":    obj.Coordinates.Point,
						"devId": dev.ID,
						"objId": obj.ID,
					}).Error("failed to build routes")
					continue
				}

				for _, route := range routes {
					if nearestInfrastructureStats[infraType] == 0 || route.Distance < nearestInfrastructureStats[infraType] {
						nearestInfrastructureStats[infraType] = route.Distance
					}
				}
			}
		}

		stats := development.Stats{
			Distance: development.POI{
				Kindergarten: int64(nearestInfrastructureStats[osm.Kindergarten]),
				School:       int64(nearestInfrastructureStats[osm.School]),
				Hospital:     int64(nearestInfrastructureStats[osm.Hospital]),
				Shops:        int64(nearestInfrastructureStats[osm.Shops]),
				Sport:        int64(nearestInfrastructureStats[osm.Sport]),
				BusStop:      int64(nearestInfrastructureStats[osm.BusStop]),
			},
		}

		for r, poiPtr := range map[int]*development.POI{
			1000: &stats.Object1000MCounts,
			2000: &stats.Object2000MCounts,
			3000: &stats.Object3000MCounts,
			4000: &stats.Object4000MCounts,
			5000: &stats.Object5000MCounts,
		} {
			radiusInfrastructure, err := infrastructureStorage.InfrastructureRadiusBoard(ctx, int(dev.ID), r)
			if err != nil {
				log.WithContext(ctx).WithError(err).Error("failed to infrastructure radius board")
				continue
			}

			radiusInfrastructureStats := make(map[osm.ObjType]int64, 6)
			for _, infraObj := range radiusInfrastructure {
				radiusInfrastructureStats[infraObj.Type]++
			}

			*poiPtr = development.POI{
				Kindergarten: radiusInfrastructureStats[osm.Kindergarten],
				School:       radiusInfrastructureStats[osm.School],
				Hospital:     radiusInfrastructureStats[osm.Hospital],
				Shops:        radiusInfrastructureStats[osm.Shops],
				Sport:        radiusInfrastructureStats[osm.Sport],
				BusStop:      radiusInfrastructureStats[osm.BusStop],
			}
		}

		err = developmentStorage.UpdateDevelopmentStats(ctx, dev.ID, stats)
		if err != nil {
			log.WithContext(ctx).WithError(err).WithFields(map[string]any{
				"devId": dev.ID,
			}).Error("failed to update development stats")
		}
	}
}
