package infrastructure

import (
	"context"
	"fmt"
	"github.com/paulmach/orb"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm"
)

func (s *Storage) NearestInfrastructure(ctx context.Context, point orb.Point, count int) (map[osm.ObjType][]Obj, error) {
	const query = `
		WITH RankedObjects AS (
			SELECT
				osm_id,
				name,
				way,
				CASE
					WHEN tags ? 'hospital'     THEN 'hospital'
					WHEN tags ? 'sport'        THEN 'sport'
					WHEN tags ? 'shops'        THEN 'shops'
					WHEN tags ? 'kindergarten' THEN 'kindergarten'
					WHEN tags ? 'bus_stop'     THEN 'bus_stop'
					WHEN tags ? 'school'       THEN 'school'
				END AS object_type,
				ROW_NUMBER() OVER (
					PARTITION BY
						CASE
							WHEN tags ? 'hospital'     THEN 'hospital'
							WHEN tags ? 'sport'        THEN 'sport'
							WHEN tags ? 'shops'        THEN 'shops'
							WHEN tags ? 'kindergarten' THEN 'kindergarten'
							WHEN tags ? 'bus_stop'     THEN 'bus_stop'
							WHEN tags ? 'school'       THEN 'school'
						END
					ORDER BY
						way <-> ST_SetSRID(ST_MakePoint($1, $2), 4326)
				) AS rn
			FROM
				osm_node
			WHERE
				tags ?| ARRAY['hospital', 'sport', 'shops', 'kindergarten', 'bus_stop', 'school']
		)
		SELECT
			osm_id AS id,
			name,
			object_type AS type,
			ST_AsGeoJSON(way) AS coordinates
		FROM
			RankedObjects
		WHERE
			rn <= $3
		ORDER BY
			object_type, rn;
	`

	rows, err := s.pg.Query(ctx, query, point.Lon(), point.Lat(), count)
	if err != nil {
		return nil, fmt.Errorf("query nearest infrastructure: %w", err)
	}
	defer rows.Close()

	nodes := make(map[osm.ObjType][]Obj, 6)
	for rows.Next() {
		var node Obj
		if err := rows.Scan(&node.ID, &node.Name, &node.Type, &node.Coordinates); err != nil {
			return nil, fmt.Errorf("scan nearest infrastructure row: %w", err)
		}
		nodes[node.Type] = append(nodes[node.Type], node)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration error: %w", rows.Err())
	}

	return nodes, nil
}
