package development

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/paulmach/orb"
)

func (s *Storage) SearchDevelopmentByBoard(ctx context.Context, bottomRight, topLeft orb.Point) ([]Development, error) {
	st := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	b := st.Select("id", "name", "ST_AsGeoJSON(cords) as coordinates, created_at").
		From("development").
		Where("ST_Contains(ST_MakeEnvelope($1, $2, $3, $4, 4326), cords)",
			bottomRight.Lon(), bottomRight.Lat(),
			topLeft.Lon(), topLeft.Lat(),
		).Where(squirrel.Eq{"deleted_at": nil})

	sql, args, err := b.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building SQL: %w", err)
	}

	rows, err := s.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var developments []Development
	for rows.Next() {
		var dev Development
		if err := rows.Scan(&dev.ID, &dev.Name, &dev.Coordinates, &dev.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		developments = append(developments, dev)
	}

	return developments, nil

}
