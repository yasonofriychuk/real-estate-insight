package selection

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *Storage) CreateSelection(ctx context.Context, profileId uuid.UUID, selection SelectionCreate) (uuid.UUID, error) {
	query := `
		INSERT INTO selection 
		    (id, profile_id, name, comment, form) 
		VALUES 
		    ($1, $2, $3, $4, $5)
	`
	selectionId := uuid.New()
	if _, err := s.pg.Exec(ctx, query, selectionId, profileId, selection.Name, selection.Comment, selection.Form); err != nil {
		return uuid.Nil, fmt.Errorf("pg.Exec: %w", err)
	}

	return selectionId, nil
}
