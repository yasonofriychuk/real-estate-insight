package objects_find_nearest_infrastructure

import (
	"context"

	"github.com/paulmach/orb"

	persistence_osm "github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/osm"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

type osmStorage interface {
	GetNearestInfrastructure(ctx context.Context, point orb.Point, objTypes []persistence_osm.ObjType) ([]persistence_osm.Obj, error)
}

type routesBuilder interface {
	BuildRoute(from, to orb.Point, transportType route_builder.TransportType) ([]route_builder.Route, error)
}
