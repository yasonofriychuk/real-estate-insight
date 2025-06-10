package infrastructure_heatmap

import (
	"context"
	"encoding/json"
	"github.com/paulmach/orb"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"math"
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/infrastructure"
)

type HeatmapHandler struct {
	log              logger.Log
	storage          infrastructureStorage
	selectionStorage selectionStorage
}

func New(log logger.Log, storage infrastructureStorage, selectionStorage selectionStorage) *HeatmapHandler {
	return &HeatmapHandler{
		log:              log,
		storage:          storage,
		selectionStorage: selectionStorage,
	}
}

func (h *HeatmapHandler) GenerateInfrastructureHeatmap(
	ctx context.Context,
	request *api.GenerateInfrastructureHeatmapReq,
) (api.GenerateInfrastructureHeatmapRes, error) {
	weights := infrastructure.HeatmapParamsWeights{
		Hospital:     1,
		Sport:        1,
		Shops:        1,
		Kindergarten: 1,
		BusStop:      1,
		School:       1,
	}

	if request.SelectionId.Set {
		selection, err := h.selectionStorage.GetById(ctx, request.SelectionId.Value, pointer.Get(auth.ProfileIdFromCtx(ctx)))
		if err != nil {
			h.log.WithContext(ctx).WithError(err).WithFields(map[string]any{
				"selection_id": request.SelectionId.Value,
				"profile_id":   pointer.Get(auth.ProfileIdFromCtx(ctx)),
			}).Error("failed get selection by id")

			return pointer.To(api.GenerateInfrastructureHeatmapInternalServerError(
				errors.BuildError(http.StatusInternalServerError, "internal error"),
			)), nil
		}
		if selection != nil {
			weights = infrastructure.HeatmapParamsWeights{
				Hospital:     selection.Form.WHospital,
				Sport:        selection.Form.WSport,
				Shops:        selection.Form.WShop,
				Kindergarten: selection.Form.WKindergarten,
				BusStop:      selection.Form.WBusStop,
				School:       selection.Form.WSchool,
			}
		}
	}

	topLeft := orb.Point{request.Bbox.TopLeftLon, request.Bbox.TopLeftLat}
	bottomRight := orb.Point{request.Bbox.BottomRightLon, request.Bbox.BottomRightLat}

	widthMeters := haversine(
		orb.Point{topLeft[0], topLeft[1]},
		orb.Point{bottomRight[0], topLeft[1]},
	)

	if widthMeters > 86000 {
		return pointer.To[api.GenerateInfrastructureHeatmapOKApplicationJSON]([]api.GenerateInfrastructureHeatmapOKItem{}), nil
	}

	params := infrastructure.HeatmapParams{
		BBox: infrastructure.HeatmapParamsBBox{
			TopLeftLon:     request.Bbox.TopLeftLon,
			TopLeftLat:     request.Bbox.TopLeftLat,
			BottomRightLon: request.Bbox.BottomRightLon,
			BottomRightLat: request.Bbox.BottomRightLat,
		},
		Weights:  weights,
		CellSize: int(max((widthMeters/15)-float64(int64(widthMeters/15)%100), 50)),
	}

	hexagons, err := h.storage.InfrastructureHeatmap(ctx, params)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed to generate infrastructure heatmap")
		return pointer.To(api.GenerateInfrastructureHeatmapInternalServerError(
			errors.BuildError(http.StatusInternalServerError, "internal error"),
		)), nil
	}

	result := make([]api.GenerateInfrastructureHeatmapOKItem, 0, len(hexagons))

	for _, hex := range hexagons {
		geometryBytes, err := json.Marshal(hex.Geometry)
		if err != nil {
			h.log.WithContext(ctx).WithError(err).Error("failed to marshal geometry")
			continue
		}

		var geom api.GenerateInfrastructureHeatmapOKItemGeometry
		if err := geom.UnmarshalJSON(geometryBytes); err != nil {
			h.log.WithContext(ctx).WithError(err).Error("failed to unmarshal geometry into API type")
			continue
		}

		result = append(result, api.GenerateInfrastructureHeatmapOKItem{
			Geometry:    geom,
			TotalWeight: hex.TotalWeight,
		})
	}

	return pointer.To[api.GenerateInfrastructureHeatmapOKApplicationJSON](result), nil
}

// Формула гаверсинуса для расчета расстояния между двумя координатами
func haversine(p1, p2 orb.Point) float64 {
	const R = 6371000 // Радиус Земли в метрах

	lat1 := toRadians(p1[1])
	lat2 := toRadians(p2[1])
	dLat := toRadians(p2[1] - p1[1])
	dLon := toRadians(p2[0] - p1[0])

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func toRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
