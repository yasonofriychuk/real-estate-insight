package infrastructure_radius_board

import (
	"context"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/infrastructure"
)

type infrastructureStorage interface {
	InfrastructureRadiusBoard(ctx context.Context, id, radius int) ([]infrastructure.Obj, error)
}
