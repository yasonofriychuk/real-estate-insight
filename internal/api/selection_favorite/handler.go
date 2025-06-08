package selection_favorite

import (
	"context"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
)

type SelectionFavoriteHandler struct {
	log logger.Log
}

func New(log logger.Log) *SelectionFavoriteHandler {
	return &SelectionFavoriteHandler{
		log: log,
	}
}

func (h *SelectionFavoriteHandler) AddToFavoriteSelection(ctx context.Context, request *api.AddToFavoriteSelectionReq) (api.AddToFavoriteSelectionRes, error) {
	return &api.AddToFavoriteSelectionOK{Status: true}, nil
}
