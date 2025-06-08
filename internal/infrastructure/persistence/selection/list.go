package selection

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *Storage) List(ctx context.Context, profileId uuid.UUID) ([]Selection, error) {
	query := `
		SELECT 
		    id, name, comment, form, created_at, updated_at 
		FROM selection 
		WHERE profile_id = $1
	`

	rows, err := s.pg.Query(ctx, query, profileId)
	if err != nil {
		return nil, fmt.Errorf("pg.Query: %w", err)
	}
	defer rows.Close()

	var selections []Selection

	for rows.Next() {
		var sel Selection

		err := rows.Scan(
			&sel.Id,
			&sel.Name,
			&sel.Comment,
			&sel.Form,
			&sel.CreatedAt,
			&sel.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		selections = append(selections, sel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return selections, nil
}
