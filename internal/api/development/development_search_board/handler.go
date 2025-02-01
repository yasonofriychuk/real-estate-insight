package development_search_board

import (
	"context"
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/paulmach/orb"
	"github.com/samber/lo"

	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
)

type DevelopmentSearchBoardHandler struct {
	log                logger.Log
	developmentStorage developmentStorage
}

func New(log logger.Log, developmentStorage developmentStorage) *DevelopmentSearchBoardHandler {
	return &DevelopmentSearchBoardHandler{
		log:                log,
		developmentStorage: developmentStorage,
	}
}

func (h *DevelopmentSearchBoardHandler) DevelopmentSearchBoard(ctx context.Context, params api.DevelopmentSearchBoardParams) (api.DevelopmentSearchBoardRes, error) {
	devs, err := h.developmentStorage.SearchDevelopmentByBoard(
		ctx,
		orb.Point{params.BottomRightLon, params.BottomRightLat},
		orb.Point{params.TopLeftLon, params.TopLeftLat},
	)

	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("development search board")
		return pointer.To(
			api.DevelopmentSearchBoardInternalServerError(
				errors.BuildError(http.StatusInternalServerError, err.Error()),
			),
		), err
	}

	response := lo.Map(devs, func(dev development.Development, _ int) api.DevelopmentSearchBoardOKItem {
		return api.DevelopmentSearchBoardOKItem{
			ID:   dev.ID,
			Name: dev.Name,
			Coords: api.OptDevelopmentSearchBoardOKItemCoords{
				Value: api.DevelopmentSearchBoardOKItemCoords{
					Lon: dev.Coordinates.Lon(),
					Lat: dev.Coordinates.Lat(),
				},
				Set: true,
			},
		}
	})

	return pointer.To[api.DevelopmentSearchBoardOKApplicationJSON](response), nil
}
