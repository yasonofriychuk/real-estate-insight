package selection_create

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
	"net/http"
)

type SelectionCreateHandler struct {
	log     logger.Log
	storage storage
}

func New(log logger.Log, storage storage) *SelectionCreateHandler {
	return &SelectionCreateHandler{
		log:     log,
		storage: storage,
	}
}

func (h *SelectionCreateHandler) CreateSelection(ctx context.Context, request *api.CreateSelectionReq) (api.CreateSelectionRes, error) {
	selectionId, err := h.storage.CreateSelection(ctx, pointer.Get(auth.ProfileIdFromCtx(ctx)), selection.SelectionCreate{
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
		h.log.WithContext(ctx).WithError(err).Error("create selection failed")
		return pointer.To(api.CreateSelectionInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	return &api.CreateSelectionOK{SelectionId: selectionId}, nil
}
