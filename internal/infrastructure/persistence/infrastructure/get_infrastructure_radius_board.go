package infrastructure

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) InfrastructureRadiusBoard(ctx context.Context, id, radius int) ([]Obj, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select(
			"osm_id",
			"name",
			"(SELECT unnest(akeys(tags)) LIMIT 1) AS key",
			"ST_AsGeoJSON(way) AS coordinates",
		).
		From("osm_node").
		Where(squirrel.Expr(
			"ST_DWithin(ST_Transform(way::geometry, 3857), ST_Transform(ST_SetSRID((?), 4326), 3857), ?)",
			squirrel.Select("cords").
				From("development").
				Where(squirrel.Eq{"id": id}),
			radius,
		)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("sql build: %w", err)
	}

	rows, err := s.pg.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	nodes := make([]Obj, 0, 100)
	for rows.Next() {
		var node Obj
		if err := rows.Scan(&node.ID, &node.Name, &node.Type, &node.Coordinates); err != nil {
			return nil, fmt.Errorf("rows: %w", err)
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
