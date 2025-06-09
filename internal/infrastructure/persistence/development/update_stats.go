package development

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
)

func (s *Storage) UpdateDevelopmentStats(ctx context.Context, developmentId int64, stats Stats) error {
	metaJSON := map[string]interface{}{
		"stats": stats,
	}

	metaBytes, err := json.Marshal(metaJSON)
	if err != nil {
		return fmt.Errorf("failed to marshal stats for development_id %d: %w", developmentId, err)
	}

	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Update("development").
		Set("meta", squirrel.Expr("meta || ?::jsonb", string(metaBytes))).
		Where(squirrel.Eq{"id": developmentId}).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to build SQL query for development_id %d: %w", developmentId, err)
	}

	_, err = s.pg.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute update query for development_id %d: %w", developmentId, err)
	}

	return nil
}
