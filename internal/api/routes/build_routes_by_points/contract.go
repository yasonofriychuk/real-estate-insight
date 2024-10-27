package build_routes_by_points

import (
	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

type routesBuilder interface {
	BuildRoute(from, to orb.Point, transportType route_builder.TransportType) ([]route_builder.Route, error)
}
