package build_routes_by_points

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AlekSi/pointer"
	geojson "github.com/paulmach/go.geojson"
	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

type BuildRoutesByPointsHandler struct {
	log           logger.Log
	routesBuilder routesBuilder
}

func New(log logger.Log, routesBuilder routesBuilder) *BuildRoutesByPointsHandler {
	return &BuildRoutesByPointsHandler{
		log:           log,
		routesBuilder: routesBuilder,
	}
}

func (h *BuildRoutesByPointsHandler) BuildRoutesByPoints(ctx context.Context, params api.BuildRoutesByPointsParams) (api.BuildRoutesByPointsRes, error) {
	fromPoint := orb.Point{params.LonFrom.Value, params.LatFrom.Value}
	toPoint := orb.Point{params.LonTo.Value, params.LatTo.Value}

	routes, err := h.routesBuilder.BuildRoute(fromPoint, toPoint, route_builder.FootTransportType)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed to build routes")
		return pointer.To(
			api.BuildRoutesByPointsInternalServerError(
				errors.BuildError(http.StatusInternalServerError, "failed to build routes"),
			),
		), err
	}

	featureCollection := geojson.NewFeatureCollection()
	for _, route := range routes {
		feature := geojson.NewLineStringFeature(route.Coordinates)
		feature.SetProperty("distance", fmt.Sprintf("%.2f meters", route.Distance))
		featureCollection.AddFeature(feature)
	}

	b, err := featureCollection.MarshalJSON()
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed to marshal featureCollection")
		return pointer.To(
			api.BuildRoutesByPointsInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), err
	}

	res := make(api.BuildRoutesByPointsOK)
	if err := res.UnmarshalJSON(b); err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed res.UnmarshalJSON(b)")
		return pointer.To(
			api.BuildRoutesByPointsInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), err
	}

	return &res, nil
}
