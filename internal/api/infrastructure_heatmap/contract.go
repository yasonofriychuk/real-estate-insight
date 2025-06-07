package infrastructure_heatmap

import (
	"context"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/infrastructure"
)

type storage interface {
	InfrastructureHeatmap(ctx context.Context, in infrastructure.HeatmapParams) ([]infrastructure.HexagonWeight, error)
}
