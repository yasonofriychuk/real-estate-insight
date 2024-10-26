package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	geojson "github.com/paulmach/go.geojson"
	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/interanal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/interanal/osm/route_builder"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger(slog.LevelDebug, "dev", os.Stdout)
	rb := route_builder.NewRouteBuilder()

	p1 := orb.Point{73.45854604143702, 61.25336620862143}
	p2 := orb.Point{73.34854329623607, 61.281014899219855}

	featureCollection := geojson.NewFeatureCollection()

	for _, t := range route_builder.TransportTypes[:1] {
		routes, err := rb.BuildRoute(p1, p2, t)
		if err != nil {
			log.WithContext(ctx).WithError(err).Error("failed to build route")
			os.Exit(1)
		}

		for _, route := range routes {
			feature := geojson.NewLineStringFeature(route.Coordinates)

			feature.SetProperty("duration", fmt.Sprintf("%.2f minutes", float64(route.Duration)/60.0))
			feature.SetProperty("distance", fmt.Sprintf("%.2f meters", route.Distance))

			featureCollection.AddFeature(feature)
		}
	}

	b, err := featureCollection.MarshalJSON()
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("failed to marshal featureCollection")
		os.Exit(1)
	}

	fmt.Println(string(b))
}
