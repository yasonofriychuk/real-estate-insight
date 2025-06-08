package selection_by_id

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/samber/lo"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
	"net/http"
)

const defaultImageUrl = "/images/no-image.jpeg"

type SelectionByIdHandler struct {
	log                logger.Log
	selectionStorage   selectionStorage
	developmentStorage developmentStorage
}

func New(log logger.Log, selectionStorage selectionStorage, developmentStorage developmentStorage) *SelectionByIdHandler {
	return &SelectionByIdHandler{
		log:                log,
		selectionStorage:   selectionStorage,
		developmentStorage: developmentStorage,
	}
}

func (h *SelectionByIdHandler) SelectionById(ctx context.Context, params api.SelectionByIdParams) (api.SelectionByIdRes, error) {
	selectionId := params.SelectionId
	profileId := pointer.Get(auth.ProfileIdFromCtx(ctx))

	s, err := h.selectionStorage.GetById(ctx, selectionId, profileId)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("get selection by id failed")
		return pointer.To(api.SelectionByIdInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	if s == nil {
		return pointer.To(api.SelectionByIdNotFound(errors.BuildError(
			http.StatusNotFound, "selection not found",
		))), nil
	}

	favoriteDevs, err := h.selectionStorage.FavoriteDevelopments(ctx, selectionId, profileId)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("get favorite developments")
		return pointer.To(api.SelectionByIdInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	developments, err := h.developmentStorage.SearchDevelopmentByFilters(ctx, development.Filter{
		DevelopmentIds: pointer.To(lo.MapToSlice(favoriteDevs, func(key int64, _ bool) int64 {
			return key
		})),
	}, nil)

	return &api.SelectionByIdOK{
		Selection: api.Selection{
			SelectionId: s.Id,
			Name:        s.Name,
			Comment:     s.Comment,
			CreatedAt:   s.CreatedAt,
			UpdatedAt:   s.UpdatedAt,
			Form: api.SelectionForm{
				LocationID:    s.Form.LocationID,
				WHospital:     s.Form.WHospital,
				WSport:        s.Form.WSport,
				WShop:         s.Form.WShop,
				WKindergarten: s.Form.WKindergarten,
				WBusStop:      s.Form.WBusStop,
				WSchool:       s.Form.WSchool,
			},
		},
		FavoriteDevelopments: lo.Map(developments, func(dev development.Development, _ int) api.SelectionByIdOKFavoriteDevelopmentsItem {
			return api.SelectionByIdOKFavoriteDevelopmentsItem{
				Development: api.Development{
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
					IsFavorite:  true,
				},
				Object3000mCounts: api.SelectionByIdOKFavoriteDevelopmentsItemObject3000mCounts{
					Kindergarten: 1,
					School:       1,
					Hospital:     1,
					Shops:        10,
					Sport:        2,
					BusStop:      5,
				},
				Distance: api.SelectionByIdOKFavoriteDevelopmentsItemDistance{
					Kindergarten: 1000,
					School:       1000,
					Hospital:     3400,
					Shops:        5444,
					Sport:        1111,
					BusStop:      1111,
				},
			}
		}),
	}, nil
}
