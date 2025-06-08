package location_list

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/samber/lo"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/location"
	"net/http"
)

type LocationListHandler struct {
	log     logger.Log
	storage storage
}

func New(log logger.Log, storage storage) *LocationListHandler {
	return &LocationListHandler{
		log:     log,
		storage: storage,
	}
}

func (h *LocationListHandler) LocationList(ctx context.Context) (api.LocationListRes, error) {
	locations, err := h.storage.CityList(ctx)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("get locationList failed")
		return pointer.To(api.LocationListInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	return &api.LocationListOK{
		Locations: lo.Map(locations, func(loc location.Location, _ int) api.LocationListOKLocationsItem {
			return api.LocationListOKLocationsItem{
				LocationId: int(loc.Id),
				Name:       loc.Name,
			}
		}),
	}, nil
}
