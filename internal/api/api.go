package api

import (
	"github.com/yasonofriychuk/real-estate-insight/internal/api/build_routes_by_points"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/development_search_filter"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/infrastructure_heatmap"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/infrastructure_radius_board"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/location_list"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/profile_login"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_create"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_delete"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_edit"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_favorite"
	"github.com/yasonofriychuk/real-estate-insight/internal/api/selection_list"
)

type API struct {
	*build_routes_by_points.BuildRoutesByPointsHandler
	*development_search_filter.DevelopmentSearchHandler
	*infrastructure_radius_board.InfrastructureRadiusBoardHandler
	*infrastructure_heatmap.HeatmapHandler
	*profile_login.ProfileLoginHandler
	*selection_delete.SelectionDeleteHandler
	*selection_favorite.SelectionFavoriteHandler
	*selection_create.SelectionCreateHandler
	*selection_list.SelectionListHandler
	*location_list.LocationListHandler
	*selection_edit.SelectionEditHandler
}
