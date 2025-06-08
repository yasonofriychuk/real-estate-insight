package selection_delete

import (
	"context"
	"github.com/google/uuid"
)

type storage interface {
	Delete(ctx context.Context, profileId, selectionId uuid.UUID) error
}
