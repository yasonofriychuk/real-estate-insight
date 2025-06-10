package infrastructure_heatmap

import (
	"context"
	"github.com/google/uuid"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/infrastructure"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
)

type infrastructureStorage interface {
	InfrastructureHeatmap(ctx context.Context, in infrastructure.HeatmapParams) ([]infrastructure.HexagonWeight, error)
}

type selectionStorage interface {
	GetById(ctx context.Context, selectionId, profileId uuid.UUID) (*selection.Selection, error)
}
