package index_page_handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
)

type IndexPageHandler struct {
	log logger.Log
}

func New(log logger.Log) *IndexPageHandler {
	return &IndexPageHandler{log: log}
}

func (h IndexPageHandler) IndexPage(ctx context.Context) (api.IndexPageRes, error) {
	f, err := os.Open("static/index.html")
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("failed create template")
		return nil, fmt.Errorf(http.StatusText(http.StatusInternalServerError))
	}
	return &api.IndexPageOK{Data: f}, nil
}
