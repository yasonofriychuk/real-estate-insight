package selection_by_id

import (
	"context"
	"github.com/google/uuid"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
)

type selectionStorage interface {
	GetById(ctx context.Context, selectionId, profileId uuid.UUID) (*selection.Selection, error)
	FavoriteDevelopments(ctx context.Context, selectionId, profileId uuid.UUID) (map[int64]bool, error)
}

type developmentStorage interface {
	SearchDevelopmentByFilters(ctx context.Context, filter development.Filter, pagination *development.Pagination) ([]development.Development, error)
}
