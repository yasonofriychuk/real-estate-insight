package selection_edit

import (
	"context"
	"github.com/google/uuid"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/selection"
)

type storage interface {
	EditSelection(ctx context.Context, profileId, selectionId uuid.UUID, updated selection.SelectionCreate) error
}
