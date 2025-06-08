package selection_edit

import (
	"context"
	stderrors "errors"
	"github.com/AlekSi/pointer"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
	"net/http"
)

type SelectionEditHandler struct {
	log     logger.Log
	storage storage
}

func New(log logger.Log, storage storage) *SelectionEditHandler {
	return &SelectionEditHandler{log: log, storage: storage}
}

func (h *SelectionEditHandler) EditSelection(ctx context.Context, request *api.EditSelectionReq) (api.EditSelectionRes, error) {
	profileId := pointer.Get(auth.ProfileIdFromCtx(ctx))
	selectionId := request.SelectionId

	err := h.storage.EditSelection(ctx, profileId, selectionId, selection.SelectionCreate{
		Name:    request.Name,
		Comment: request.Comment,
		Form: selection.Form{
			LocationID:    request.Form.LocationID,
			WHospital:     request.Form.WHospital,
			WSport:        request.Form.WSport,
			WShop:         request.Form.WShop,
			WKindergarten: request.Form.WKindergarten,
			WBusStop:      request.Form.WBusStop,
			WSchool:       request.Form.WSchool,
		},
	})
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("edit selection failed")
		if stderrors.Is(err, selection.ErrAccessDeniedOrNotFound) {
			return pointer.To(api.EditSelectionNotFound(errors.BuildError(
				http.StatusNotFound, "selection not found",
			))), nil
		}

		return pointer.To(api.EditSelectionInternalServerError(errors.BuildError(
			http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError),
		))), nil
	}

	return &api.EditSelectionOK{Status: true}, nil
}
