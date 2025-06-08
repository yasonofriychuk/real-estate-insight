package selection

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *Storage) FavoriteDevelopments(ctx context.Context, selectionId, profileId uuid.UUID) (map[int64]bool, error) {
	query := `
		SELECT fs.development_id 
		FROM favorite_selection_development fs
		WHERE 
			selection_id = $1 AND
			exists(SELECT 1 FROM selection s WHERE s.id = $1 AND s.profile_id = $2)
	`

	rows, err := s.pg.Query(ctx, query, selectionId, profileId)
	if err != nil {
		return nil, fmt.Errorf("pg.Query: %w", err)
	}
	defer rows.Close()

	favorites := make(map[int64]bool)
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		favorites[id] = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return favorites, nil
}

func (s *Storage) SetFavoriteDevelopment(ctx context.Context, developmentId int64, selectionId, profileId uuid.UUID, value bool) error {
	// Проверяем, что profileId владеет selectionId
	var ownerProfileId uuid.UUID
	query := `
		SELECT profile_id 
		FROM selection 
		WHERE id = $1
	`
	if err := s.pg.QueryRow(ctx, query, selectionId).Scan(&ownerProfileId); err != nil {
		return fmt.Errorf("failed to find selection with id %s: %w", selectionId, err)
	}

	// Если profileId не является владельцем selectionId, ничего не делаем
	if ownerProfileId != profileId {
		return nil
	}

	if value {
		query = `
			INSERT INTO favorite_selection_development (development_id, selection_id)
			VALUES ($1, $2)
			ON CONFLICT (development_id, selection_id) DO NOTHING
		`
		if _, err := s.pg.Exec(ctx, query, developmentId, selectionId); err != nil {
			return fmt.Errorf("failed to add to favorites: %w", err)
		}
	} else {
		query = `
			DELETE FROM favorite_selection_development
			WHERE development_id = $1 AND selection_id = $2
		`
		if _, err := s.pg.Exec(ctx, query, developmentId, selectionId); err != nil {
			return fmt.Errorf("failed to remove from favorites: %w", err)
		}
	}

	return nil
}
