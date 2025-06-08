package selection_list

import (
	"context"
	"github.com/google/uuid"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
)

type storage interface {
	List(ctx context.Context, profileId uuid.UUID) ([]selection.Selection, error)
}
