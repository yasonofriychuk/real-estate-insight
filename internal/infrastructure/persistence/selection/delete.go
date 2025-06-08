package selection

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *Storage) Delete(ctx context.Context, profileId, selectionId uuid.UUID) error {
	query := `
		DELETE FROM selection WHERE profile_id = $1 AND id = $2
	`
	if _, err := s.pg.Exec(ctx, query, profileId, selectionId); err != nil {
		return fmt.Errorf("pg.Exec: %w", err)
	}
	return nil
}
