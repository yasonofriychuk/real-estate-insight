package osm

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
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
		var obj Obj
		err := rows.Scan(&obj.Type, &obj.Name, &obj.Coordinates)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		objs = append(objs, obj)
	}

	return objs, nil
}

func modifyStatementByType(st squirrel.SelectBuilder, objType ObjType) squirrel.SelectBuilder {
	switch objType {
	case Hospital:
		return st.Where(squirrel.Expr("tags->'hospital' = 'yes'"))
	case Sport:
		return st.Where(squirrel.Expr("tags->'sport' = 'yes'"))
	case Shops:
		return st.Where(squirrel.Expr("tags->'shops' = 'yes'"))
	case Kindergarten:
		return st.Where(squirrel.Expr("tags->'kindergarten' = 'yes'"))
	case BusStop:
		return st.Where(squirrel.Expr("tags->'bus_stop' = 'yes'"))
	}
	return st
}
