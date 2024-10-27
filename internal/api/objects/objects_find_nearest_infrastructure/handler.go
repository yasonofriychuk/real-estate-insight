package objects_find_nearest_infrastructure

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/AlekSi/pointer"
	geojson "github.com/paulmach/go.geojson"
	"golang.org/x/sync/errgroup"

	"github.com/yasonofriychuk/real-estate-insight/internal/api/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/osm"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm/route_builder"
)

type ObjectsFindNearestInfrastructureHandler struct {
	log           logger.Log
	storage       osmStorage
	routesBuilder routesBuilder
}

func New(log logger.Log, storage osmStorage, routesBuilder routesBuilder) *ObjectsFindNearestInfrastructureHandler {
	return &ObjectsFindNearestInfrastructureHandler{
		log:           log,
		storage:       storage,
		routesBuilder: routesBuilder,
	}
}

func (h *ObjectsFindNearestInfrastructureHandler) ObjectsFindNearestInfrastructure(ctx context.Context, params api.ObjectsFindNearestInfrastructureParams) (api.ObjectsFindNearestInfrastructureRes, error) {
	point := paramsToPoint(params)
	objs, err := h.storage.GetNearestInfrastructure(ctx, point, paramsToObjTypes(params))

	if err != nil {
		h.log.WithContext(ctx).WithError(err).WithFields(map[string]any{
			"lon":   params.Lon.Value,
			"lat":   params.Lat.Value,
			"types": params.ObjectTypes,
		}).Error("failed to get objects from storage")
		return pointer.To(
			api.ObjectsFindNearestInfrastructureInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), err
	}

	errGroup, ctx := errgroup.WithContext(ctx)
	mutex := &sync.Mutex{}
	collection := geojson.NewFeatureCollection()

	origin := geojson.NewPointFeature([]float64{params.Lon.Value, params.Lat.Value})
	origin.SetProperty("marker-symbol", "circle-stroked")
	collection.AddFeature(origin)

	for _, obj := range objs {
		f := geojson.NewPointFeature([]float64{obj.Coordinates.Lon(), obj.Coordinates.Lat()})
		f.SetProperty("objType", obj.Type)
		f.SetProperty("name", obj.Name)
		f.SetProperty("marker-color", colorByType(obj.Type))
		collection.AddFeature(f)

		errGroup.Go(func() error {
			routes, err := h.routesBuilder.BuildRoute(point, obj.Coordinates, route_builder.CarTransportType)
			if err != nil {
				return fmt.Errorf("build route: %w", err)
			}

			for _, route := range routes {
				fRoute := geojson.NewLineStringFeature(route.Coordinates)

				fRoute.SetProperty("distance", fmt.Sprintf("%.2f meters", route.Distance))
				fRoute.SetProperty("objType", obj.Type)
				fRoute.SetProperty("name", obj.Name)
				fRoute.SetProperty("stroke-width", 4.5)
				fRoute.SetProperty("stroke", colorByType(obj.Type))

				mutex.Lock()
				collection.AddFeature(fRoute)
				mutex.Unlock()
			}

			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed to build routes")
		return pointer.To(
			api.ObjectsFindNearestInfrastructureInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), err
	}

	b, err := collection.MarshalJSON()
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed to marshal featureCollection")
		return pointer.To(
			api.ObjectsFindNearestInfrastructureInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), err
	}

	res := make(api.ObjectsFindNearestInfrastructureOK)
	if err := res.UnmarshalJSON(b); err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed res.UnmarshalJSON(b)")
		return pointer.To(
			api.ObjectsFindNearestInfrastructureInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), err
	}

	return &res, nil
}

func colorByType(objType osm.ObjType) string {
	switch objType {
	case osm.Hospital:
		return "#4CAF50" // Зеленый
	case osm.Sport:
		return "#FF5722" // Оранжевый
	case osm.Shops:
		return "#FFC107" // Желтый
	case osm.Kindergarten:
		return "#FFEB3B" // Светло-желтый
	case osm.BusStop:
		return "#2196F3" // Синий
	}
	return "#9E9E9E" // Серый цвет
}
