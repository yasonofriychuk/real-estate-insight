package location

import (
	"context"
	"fmt"
)

func (s *Storage) CityList(ctx context.Context) ([]Location, error) {
	query := `
		SELECT 
		    id, name
		FROM location 
		WHERE 
		    loc_type = 'city' AND 
		    id IN (SELECT DISTINCT location_id FROM development)
	`

	rows, err := s.pg.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("pg.Query: %w", err)
	}
	defer rows.Close()

	var locations []Location

	for rows.Next() {
		var loc Location

		if err := rows.Scan(
			&loc.Id,
			&loc.Name,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		locations = append(locations, loc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return locations, nil
}
