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
	"sync"
)

type LocationListHandler struct {
	log           logger.Log
	storage       storage
	cacheLocation []api.LocationListOKLocationsItem
	mutex         sync.Mutex
}

func New(log logger.Log, storage storage) *LocationListHandler {
	return &LocationListHandler{
		log:     log,
		storage: storage,
	}
}

func (h *LocationListHandler) LocationList(ctx context.Context) (api.LocationListRes, error) {
	if h.cacheLocation != nil {
		return &api.LocationListOK{
			Locations: h.cacheLocation,
		}, nil
	}

	locations, err := h.storage.CityList(ctx)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("get locationList failed")
		return pointer.To(api.LocationListInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	apiLocations := lo.Map(locations, func(loc location.Location, _ int) api.LocationListOKLocationsItem {
		return api.LocationListOKLocationsItem{
			LocationId: int(loc.Id),
			Name:       loc.Name,
		}
	})

	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.cacheLocation = apiLocations

	return &api.LocationListOK{
		Locations: apiLocations,
	}, nil
}
