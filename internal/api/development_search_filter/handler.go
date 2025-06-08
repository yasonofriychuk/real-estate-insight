package development_search_filter

import (
	"context"
	"fmt"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"net/http"

	"github.com/AlekSi/pointer"
	"github.com/paulmach/orb"
	"github.com/samber/lo"
	"golang.org/x/sync/errgroup"

	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
)

const defaultImageUrl = "/images/no-image.jpeg"

type DevelopmentSearchHandler struct {
	log                logger.Log
	developmentStorage developmentStorage
	selectionStorage   selectionStorage
}

func New(log logger.Log, developmentStorage developmentStorage, selectionStorage selectionStorage) *DevelopmentSearchHandler {
	return &DevelopmentSearchHandler{
		log:                log,
		developmentStorage: developmentStorage,
		selectionStorage:   selectionStorage,
	}
}

func (h *DevelopmentSearchHandler) DevelopmentSearch(ctx context.Context, req *api.DevelopmentSearchReq) (api.DevelopmentSearchRes, error) {
	var filter development.Filter

	if v, ok := req.SearchQuery.Get(); ok {
		filter.SearchQuery = v
	}

	if v, ok := req.Board.Get(); ok {
		filter.Board = &development.Board{
			BottomRight: orb.Point{v.BottomRightLon, v.BottomRightLat},
			TopLeft:     orb.Point{v.TopLeftLon, v.TopLeftLat},
		}
	}

	errGroup, groupCtx := errgroup.WithContext(ctx)

	var count int64
	errGroup.Go(func() (err error) {
		count, err = h.developmentStorage.CountsDevelopmentByFilters(groupCtx, filter)
		if err != nil {
			return fmt.Errorf("count development by filters: %w", err)
		}
		return nil
	})

	var devs []development.Development
	errGroup.Go(func() (err error) {
		devs, err = h.developmentStorage.SearchDevelopmentByFilters(groupCtx, filter, nil)
		if err != nil {
			return fmt.Errorf("search development by filters: %w", err)
		}
		return nil
	})

	var favoritesDevIds map[int64]bool
	errGroup.Go(func() (err error) {
		if !req.SelectionId.Set {
			return nil
		}
		favoritesDevIds, err = h.selectionStorage.FavoriteDevelopments(ctx, req.SelectionId.Value, pointer.Get(auth.ProfileIdFromCtx(ctx)))
		if err != nil {
			return fmt.Errorf("get favorite developments: %w", err)
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		h.log.WithContext(ctx).WithError(err).Error("development search")
		return pointer.To(
			api.DevelopmentSearchInternalServerError(
				errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
			),
		), nil
	}

	response := api.DevelopmentSearchOK{
		Developments: lo.Map(devs, func(dev development.Development, _ int) api.Development {
			return api.Development{
				ID:   dev.ID,
				Name: dev.Name,
				Coords: api.DevelopmentCoords{
					Lat: dev.Coordinates.Lat(),
					Lon: dev.Coordinates.Lon(),
				},
				ImageUrl:    lo.If(dev.Meta.ImageURL != "", dev.Meta.ImageURL).Else(defaultImageUrl),
				Description: dev.Meta.Description,
				AvitoUrl:    dev.Meta.AvitoUrl,
				GisUrl:      dev.Meta.GisUrl,
				Address:     dev.Meta.Address,
				IsFavorite:  favoritesDevIds[dev.ID],
			}
		}),
		Meta: api.DevelopmentSearchOKMeta{
			Total: count,
		},
	}

	return pointer.To[api.DevelopmentSearchOK](response), nil
}
