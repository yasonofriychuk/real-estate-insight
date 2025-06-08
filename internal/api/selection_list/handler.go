package selection_list

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/samber/lo"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
	"net/http"
)

type SelectionListHandler struct {
	log     logger.Log
	storage storage
}

func New(log logger.Log, storage storage) *SelectionListHandler {
	return &SelectionListHandler{
		log:     log,
		storage: storage,
	}
}

func (h *SelectionListHandler) SelectionList(ctx context.Context) (api.SelectionListRes, error) {
	selections, err := h.storage.List(ctx, pointer.Get(auth.ProfileIdFromCtx(ctx)))
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("get selection list failed")
		return pointer.To(api.SelectionListInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	return &api.SelectionListOK{
		Selections: lo.Map(selections, func(s selection.Selection, _ int) api.SelectionListOKSelectionsItem {
			return api.SelectionListOKSelectionsItem{
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
			}
		}),
	}, nil
}
