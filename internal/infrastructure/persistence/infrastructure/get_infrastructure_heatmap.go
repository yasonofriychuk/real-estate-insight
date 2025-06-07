package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Storage) InfrastructureHeatmap(ctx context.Context, in HeatmapParams) ([]HexagonWeight, error) {
	const query = `
		WITH bbox AS (
			SELECT ST_Transform(ST_MakeEnvelope($1, $2, $3, $4, 4326), 3857) AS geom
		),
		hexgrid AS (
			SELECT (ST_HexagonGrid($5, bbox.geom)).geom AS geom FROM bbox
		),
		filtered_nodes AS (
			SELECT way,
			CASE
				WHEN tags -> 'hospital' = 'yes' THEN $6
				WHEN tags -> 'sport' = 'yes' THEN $7
				WHEN tags -> 'shops' = 'yes' THEN $8
				WHEN tags -> 'kindergarten' = 'yes' THEN $9
				WHEN tags -> 'bus_stop' = 'yes' THEN $10
				WHEN tags -> 'school' = 'yes' THEN $11
				ELSE 0
				END AS weight
			FROM osm_node
			WHERE
				way && ST_MakeEnvelope($1, $2, $3, $4, 4326)
		),
		node_geoms AS (
			SELECT ST_Transform(way, 3857) AS geom, weight FROM filtered_nodes WHERE weight > 0
		),
		aggregated AS (
			SELECT h.geom, SUM(n.weight) AS total_weight FROM hexgrid h
			LEFT JOIN node_geoms n  ON ST_Intersects(h.geom, n.geom)
			GROUP BY h.geom
			HAVING SUM(n.weight) > 0
		)
		SELECT
			ST_AsGeoJSON(ST_Transform(geom, 4326))::json AS geometry,
			total_weight
		FROM aggregated;
		`

	rows, err := s.pg.Query(ctx, query,
		in.BBox.TopLeftLon,
		in.BBox.BottomRightLat,
		in.BBox.BottomRightLon,
		in.BBox.TopLeftLat,
		in.CellSize,
		in.Weights.Hospital,
		in.Weights.Sport,
		in.Weights.Shops,
		in.Weights.Kindergarten,
		in.Weights.BusStop,
		in.Weights.School,
	)

	if err != nil {
		return nil, fmt.Errorf("query heatmap: %w", err)
	}
	defer rows.Close()
	var result []HexagonWeight
	var maxWeight float64

	for rows.Next() {
		var hw HexagonWeight
		var rawGeometry []byte

		if err := rows.Scan(&rawGeometry, &hw.TotalWeight); err != nil {
			return nil, fmt.Errorf("scan heatmap row: %w", err)
		}
		if err := json.Unmarshal(rawGeometry, &hw.Geometry); err != nil {
			return nil, fmt.Errorf("unmarshal geojson: %w", err)
		}

		if hw.TotalWeight > maxWeight {
			maxWeight = hw.TotalWeight
		}

		result = append(result, hw)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	if maxWeight > 0 {
		for i := range result {
			result[i].TotalWeight = result[i].TotalWeight / maxWeight
		}
	}

	return result, nil
}
