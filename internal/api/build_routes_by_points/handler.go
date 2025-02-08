package build_routes_by_points

import (
	"context"
	std_errors "errors"
	"fmt"
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/jackc/pgx/v5"
	geojson "github.com/paulmach/go.geojson"
	"github.com/paulmach/orb"
	"golang.org/x/sync/errgroup"

	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

type BuildRoutesByPointsHandler struct {
	log           logger.Log
	routesBuilder routesBuilder
	storage       storage
}

func New(log logger.Log, routesBuilder routesBuilder, storage storage) *BuildRoutesByPointsHandler {
	return &BuildRoutesByPointsHandler{
		log:           log,
		routesBuilder: routesBuilder,
		storage:       storage,
	}
}

func (h *BuildRoutesByPointsHandler) BuildRoutesByPoints(ctx context.Context, params api.BuildRoutesByPointsParams) (api.BuildRoutesByPointsRes, error) {
	var fromPoint, toPoint orb.Point

	errGroup, errCtx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		var err error
		fromPoint, err = h.storage.GetCoordinatesDevelopmentById(errCtx, params.DevelopmentId)
		if err != nil {
			return fmt.Errorf("storage.GetCoordinatesDevelopmentById: %w", err)
		}
		return nil
	})

	errGroup.Go(func() error {
		var err error
		toPoint, err = h.storage.GetCoordinatesOsmById(errCtx, params.OsmId)
		if err != nil {
			return fmt.Errorf("storage.GetCoordinatesOsmById: %w", err)
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		if std_errors.Is(err, pgx.ErrNoRows) {
			return pointer.To(
				api.BuildRoutesByPointsNotFound(
					errors.BuildError(http.StatusNotFound, err.Error()),
				),
			), nil
		}

		h.log.WithContext(errCtx).WithError(err).Error("failed to get points coords")
		return pointer.To(
			api.BuildRoutesByPointsInternalServerError(
				errors.BuildError(http.StatusInternalServerError, err.Error()),
			),
		), nil
	}

	routes, err := h.routesBuilder.BuildRoute(fromPoint, toPoint, route_builder.FootTransportType)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed to build routes")
		return pointer.To(
			api.BuildRoutesByPointsInternalServerError(
				errors.BuildError(http.StatusInternalServerError, "failed to build routes"),
			),
		), nil
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
		), nil
	}

	res := make(api.BuildRoutesByPointsOK)
	if err := res.UnmarshalJSON(b); err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed res.UnmarshalJSON(b)")
		return pointer.To(
			api.BuildRoutesByPointsInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), nil
	}

	return &res, nil
}
