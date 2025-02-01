package api

import (
	"github.com/yasonofriychuk/real-estate-insight/internal/api/development/development_search_board"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/objects/objects_find_nearest_infrastructure"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/routes/build_routes_by_points"
)

type API struct {
	*build_routes_by_points.BuildRoutesByPointsHandler
	*objects_find_nearest_infrastructure.ObjectsFindNearestInfrastructureHandler
	*development_search_board.DevelopmentSearchBoardHandler
}
