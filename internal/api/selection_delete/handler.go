package selection_delete

import (
	"context"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
)

type SelectionDeleteHandler struct {
	log logger.Log
}

func New(log logger.Log) *SelectionDeleteHandler {
	return &SelectionDeleteHandler{
		log: log,
	}
}

func (h *SelectionDeleteHandler) DeleteSelection(ctx context.Context, params api.DeleteSelectionParams) (api.DeleteSelectionRes, error) {
	return &api.DeleteSelectionOK{Status: true}, nil
}
