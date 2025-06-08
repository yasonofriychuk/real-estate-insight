package selection_create

import (
	"context"
	"github.com/google/uuid"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
)

type storage interface {
	CreateSelection(ctx context.Context, profileId uuid.UUID, selection selection.SelectionCreate) (uuid.UUID, error)
}
