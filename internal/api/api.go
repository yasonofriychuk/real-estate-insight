package api

import (
	"github.com/yasonofriychuk/real-estate-insight/internal/api/build_routes_by_points"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/development_search_filter"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/infrastructure_heatmap"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/infrastructure_radius_board"
)

type API struct {
	*build_routes_by_points.BuildRoutesByPointsHandler
	*development_search_filter.DevelopmentSearchHandler
	*infrastructure_radius_board.InfrastructureRadiusBoardHandler
	*infrastructure_heatmap.HeatmapHandler
}
