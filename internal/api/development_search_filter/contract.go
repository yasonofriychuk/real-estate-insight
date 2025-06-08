package development_search_filter

import (
	"context"
	"github.com/google/uuid"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/development"
)

type developmentStorage interface {
	CountsDevelopmentByFilters(ctx context.Context, filter development.Filter) (int64, error)
	SearchDevelopmentByFilters(ctx context.Context, filter development.Filter, pagination *development.Pagination) ([]development.Development, error)
}

type selectionStorage interface {
	FavoriteDevelopments(ctx context.Context, selectionId, profileId uuid.UUID) (map[int64]bool, error)
}
