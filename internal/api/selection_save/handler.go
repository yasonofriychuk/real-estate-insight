package selection_save

import (
	"context"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
)

type SelectionSaveHandler struct {
	log logger.Log
}

func New(log logger.Log) *SelectionSaveHandler {
	return &SelectionSaveHandler{
		log: log,
	}
}

func (h *SelectionSaveHandler) CreateSelection(ctx context.Context, request *api.CreateSelectionReq) (api.CreateSelectionRes, error) {
	return &api.CreateSelectionOK{Status: true}, nil
}
