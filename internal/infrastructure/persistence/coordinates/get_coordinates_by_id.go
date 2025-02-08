package coordinates

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence"
)

func (s *Storage) GetCoordinatesDevelopmentById(ctx context.Context, id int64) (orb.Point, error) {
	var p persistence.Point

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("ST_AsGeoJSON(cords) as coordinates").
		From("development").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return orb.Point{}, fmt.Errorf("sql build: %w", err)
	}
	err = s.pg.QueryRow(ctx, query, args...).Scan(&p)
	if err != nil {
		return orb.Point{}, fmt.Errorf("queryRow: %w", err)
	}

	return p.Point, nil
}

func (s *Storage) GetCoordinatesOsmById(ctx context.Context, id int64) (orb.Point, error) {
	var p persistence.Point

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("ST_AsGeoJSON(way) as coordinates").
		From("osm_node").
		Where(squirrel.Eq{"osm_id": id}).
		ToSql()
	if err != nil {
		return orb.Point{}, err
	}

	err = s.pg.QueryRow(ctx, query, args...).Scan(&p)
	if err != nil {
		return orb.Point{}, err
	}

	return p.Point, nil
}
