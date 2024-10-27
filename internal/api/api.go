package api

import (
	"github.com/yasonofriychuk/real-estate-insight/internal/api/routes/build_routes_by_points"
)

type API struct {
	*build_routes_by_points.BuildRoutesByPointsHandler
}
