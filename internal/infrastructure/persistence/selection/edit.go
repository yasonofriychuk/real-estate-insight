package selection

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var ErrAccessDeniedOrNotFound = errors.New("access denied or not found")

func (s *Storage) EditSelection(ctx context.Context, profileId, selectionId uuid.UUID, updated SelectionCreate) error {
	query := `
		UPDATE selection
		SET 
			name = $1,
			comment = $2,
			form = $3,
			updated_at = current_timestamp
		WHERE 
			id = $4 AND profile_id = $5
	`

	result, err := s.pg.Exec(ctx, query, updated.Name, updated.Comment, updated.Form, selectionId, profileId)
	if err != nil {
		return fmt.Errorf("pg.Exec (update): %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrAccessDeniedOrNotFound
	}

	return nil
}
