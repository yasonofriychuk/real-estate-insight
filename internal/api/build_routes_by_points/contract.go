package build_routes_by_points

import (
	"context"

	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

type routesBuilder interface {
	BuildRoute(from, to orb.Point, transportType route_builder.TransportType) ([]route_builder.Route, error)
}

type storage interface {
	GetCoordinatesDevelopmentById(ctx context.Context, id int64) (orb.Point, error)
	GetCoordinatesOsmById(ctx context.Context, id int64) (orb.Point, error)
}
