package selection

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func (s *Storage) GetById(ctx context.Context, selectionId, profileId uuid.UUID) (*Selection, error) {
	query := `
		SELECT 
		    id, name, comment, form, created_at, updated_at 
		FROM selection
		WHERE id = $1 AND profile_id = $2
	`

	row := s.pg.QueryRow(ctx, query, selectionId, profileId)

	var sel Selection
	err := row.Scan(
		&sel.Id,
		&sel.Name,
		&sel.Comment,
		&sel.Form,
		&sel.CreatedAt,
		&sel.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("row.Scan: %w", err)
	}
	return &sel, nil
}
