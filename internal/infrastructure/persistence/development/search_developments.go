package development

import (
	"context"
	"fmt"
	"github.com/AlekSi/pointer"

	"github.com/Masterminds/squirrel"
)

func (s *Storage) CountsDevelopmentByFilters(ctx context.Context, filter Filter) (int64, error) {
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("count(*) as count").
		From("development").
		Where(squirrel.Eq{"deleted_at": nil})

	b = modifyBuilderByFilter(b, filter)
	sql, args, err := b.ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building SQL: %w", err)
	}

	rows, err := s.pg.Query(ctx, sql, args...)
	if err != nil {
		return 0, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, fmt.Errorf("error scanning row: %w", err)
		}
	}
	return count, nil
}

func (s *Storage) SearchDevelopmentByFilters(ctx context.Context, filter Filter, pagination *Pagination) ([]Development, error) {
	b := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("id", "name", "ST_AsGeoJSON(cords) as coordinates", "created_at", "meta").
		From("development").Where(squirrel.Eq{"deleted_at": nil}).
		OrderBy("name")

	b = modifyBuilderByPagination(b, pagination)
	b = modifyBuilderByFilter(b, filter)

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
		if err := rows.Scan(&dev.ID, &dev.Name, &dev.Coordinates, &dev.CreatedAt, &dev.Meta); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		developments = append(developments, dev)
	}

	return developments, nil

}

func modifyBuilderByPagination(b squirrel.SelectBuilder, pagination *Pagination) squirrel.SelectBuilder {
	if pagination != nil {
		if pagination.PerPage > 0 {
			b = b.Limit(uint64(pagination.PerPage))
		}
		if pagination.Page > 1 {
			b = b.Offset(uint64((pagination.Page - 1) * pagination.PerPage))
		}
	}
	return b
}

func modifyBuilderByFilter(b squirrel.SelectBuilder, filter Filter) squirrel.SelectBuilder {
	if filter.Board != nil {
		b = b.Where("ST_Contains(ST_MakeEnvelope(?, ?, ?, ?, 4326), cords)",
			filter.Board.BottomRight.Lon(), filter.Board.BottomRight.Lat(),
			filter.Board.TopLeft.Lon(), filter.Board.TopLeft.Lat(),
		)
	}

	if filter.SearchQuery != "" {
		b = b.Where(squirrel.Expr("name ILIKE ?", "%"+filter.SearchQuery+"%"))
	}

	if filter.DevelopmentIds != nil {
		b = b.Where(squirrel.Eq{"id": pointer.Get(filter.DevelopmentIds)})
	}

	return b
}
