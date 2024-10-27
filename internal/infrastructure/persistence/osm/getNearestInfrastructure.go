package osm

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	geojson "github.com/paulmach/go.geojson"
	"github.com/paulmach/orb"
)

func (s *Storage) GetNearestInfrastructure(ctx context.Context, point orb.Point, objTypes []ObjType) ([]Obj, error) {
	st := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	sqls := make([]string, 0, len(objTypes))
	for _, objType := range objTypes {
		objSt := st.Select(fmt.Sprintf("'%s' as type", objType), "name", "ST_AsGeoJSON(ST_Transform(way, 4326)) AS geometry").
			From(tableName).
			OrderBy("ST_Distance(ST_Transform(way, 4326)", "(SELECT geom FROM point))").
			Where(squirrel.Expr("name is not null")).
			Limit(1)

		sql, _, err := modifyStatementByType(objSt, objType).ToSql()
		if err != nil {
			return nil, fmt.Errorf("build sql [%s]: %w", objType, err)
		}
		sqls = append(sqls, fmt.Sprintf("(%s)", sql))
	}

	query := fmt.Sprintf("WITH point AS (SELECT ST_SetSRID(ST_MakePoint($1, $2), 4326) AS geom) %s", strings.Join(sqls, " union all "))
	rows, err := s.pg.Query(ctx, query, point.Lon(), point.Lat())
	if err != nil {
		return nil, fmt.Errorf("query rows: %w", err)
	}
	defer rows.Close()

	objs := make([]Obj, 0, len(objTypes))
	for rows.Next() {
		var objectType, name, geometryData string
		err := rows.Scan(&objectType, &name, &geometryData)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		geometry, err := geojson.UnmarshalGeometry([]byte(geometryData))
		if err != nil {
			return nil, fmt.Errorf("unmarshal geometry: %w", err)
		}

		if geometry.Type != geojson.GeometryPoint {
			return nil, fmt.Errorf("invalid geometry type: %s", geometry.Type)
		}

		objs = append(objs, Obj{
			Name:        name,
			Coordinates: orb.Point(geometry.Point),
			Type:        ObjType(objectType),
		})
	}

	return objs, nil
}

func modifyStatementByType(st squirrel.SelectBuilder, objType ObjType) squirrel.SelectBuilder {
	switch objType {
	case Hospital:
		return st.Where(squirrel.Or{
			squirrel.Expr("tags->'amenity' = 'hospital'"),
			squirrel.Expr("tags->'amenity' = 'clinic'"),
		})
	case Sport:
		return st.Where(squirrel.Or{
			squirrel.Expr("tags->'leisure' = 'sports_centre'"),
			squirrel.Expr("tags->'leisure' = 'stadium'"),
			squirrel.Expr("tags->'leisure' = 'pitch'"),
			squirrel.Expr("tags->'leisure' = 'fitness_centre'"),
		})
	case Shops:
		return st.Where("tags->'shop' is not null")
	case Kindergarten:
		return st.Where(squirrel.Expr("tags->'amenity' = 'kindergarten'"))
	case BusStop:
		return st.Where(squirrel.Or{
			squirrel.Expr("tags->'amenity' = 'bus_stop'"),
			squirrel.Expr("tags->'public_transport' = 'stop_position'"),
			squirrel.Expr("tags->'public_transport' = 'station'"),
			squirrel.Expr("tags->'railway' = 'subway_entrance'"),
		})
	}
	return st
}
