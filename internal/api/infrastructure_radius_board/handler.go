package infrastructure_radius_board

import (
	"context"
	std_errors "errors"
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"

	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/infrastructure"
)

type InfrastructureRadiusBoardHandler struct {
	log                   logger.Log
	infrastructureStorage infrastructureStorage
}

func New(log logger.Log, infrastructureStorage infrastructureStorage) *InfrastructureRadiusBoardHandler {
	return &InfrastructureRadiusBoardHandler{
		log:                   log,
		infrastructureStorage: infrastructureStorage,
	}
}

func (h *InfrastructureRadiusBoardHandler) InfrastructureRadiusBoard(ctx context.Context, params api.InfrastructureRadiusBoardParams) (api.InfrastructureRadiusBoardRes, error) {
	infs, err := h.infrastructureStorage.InfrastructureRadiusBoard(ctx, params.DevelopmentId, params.Radius)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("infrastructure radius board")
		if std_errors.Is(err, pgx.ErrNoRows) {
			return pointer.To(api.InfrastructureRadiusBoardNotFound(errors.BuildError(http.StatusNotFound, err.Error()))), nil
		}
		return pointer.To(api.InfrastructureRadiusBoardInternalServerError(errors.BuildError(http.StatusInternalServerError, err.Error()))), nil
	}

	response := lo.Map(infs, func(inf infrastructure.Obj, _ int) api.InfrastructureRadiusBoardOKItem {
		return api.InfrastructureRadiusBoardOKItem{
			ID: inf.ID,
			Name: api.OptString{
				Value: inf.Name,
				Set:   inf.Name != "",
			},
			ObjType: string(inf.Type),
			Coords: api.InfrastructureRadiusBoardOKItemCoords{
				Lon: inf.Coordinates.Lon(),
				Lat: inf.Coordinates.Lat(),
			},
		}
	})

	return pointer.To[api.InfrastructureRadiusBoardOKApplicationJSON](response), nil
}
