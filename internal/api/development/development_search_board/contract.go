package development_search_board

import (
	"context"

	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
)

type developmentStorage interface {
	SearchDevelopmentByBoard(ctx context.Context, bottomRight, topLeft orb.Point) ([]development.Development, error)
}
