package selection_delete

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"net/http"
)

type SelectionDeleteHandler struct {
	log     logger.Log
	storage storage
}

func New(log logger.Log, storage storage) *SelectionDeleteHandler {
	return &SelectionDeleteHandler{
		log:     log,
		storage: storage,
	}
}

func (h *SelectionDeleteHandler) DeleteSelection(ctx context.Context, params api.DeleteSelectionParams) (api.DeleteSelectionRes, error) {
	if err := h.storage.Delete(ctx, pointer.Get(auth.ProfileIdFromCtx(ctx)), params.ID); err != nil {
		h.log.WithContext(ctx).WithError(err).Error("delete selection failed")
		return pointer.To(api.DeleteSelectionInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}
	return &api.DeleteSelectionOK{Status: true}, nil
}
