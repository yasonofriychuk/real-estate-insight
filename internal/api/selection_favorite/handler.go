package selection_favorite

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/auth"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"net/http"
)

type SelectionFavoriteHandler struct {
	log     logger.Log
	storage storage
}

func New(log logger.Log, storage storage) *SelectionFavoriteHandler {
	return &SelectionFavoriteHandler{
		log:     log,
		storage: storage,
	}
}

func (h *SelectionFavoriteHandler) AddToFavoriteSelection(ctx context.Context, request *api.AddToFavoriteSelectionReq) (api.AddToFavoriteSelectionRes, error) {
	err := h.storage.SetFavoriteDevelopment(
		ctx,
		int64(request.DevelopmentID),
		request.SelectionID,
		pointer.Get(auth.ProfileIdFromCtx(ctx)),
		request.Value,
	)

	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed to toggle favorite")
		return pointer.To(api.AddToFavoriteSelectionInternalServerError(
			errors.BuildError(http.StatusInternalServerError, "Internal Server Error"),
		)), nil
	}
	return &api.AddToFavoriteSelectionOK{Status: true}, nil
}
