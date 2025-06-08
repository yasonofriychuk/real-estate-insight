package selection_favorite

import (
	"context"
	"github.com/google/uuid"
)

type storage interface {
	SetFavoriteDevelopment(ctx context.Context, developmentId int64, selectionId, profileId uuid.UUID, value bool) error
}
