package profile

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *Storage) Login(ctx context.Context, username string, password string) (*uuid.UUID, error) {
	const query = `SELECT id FROM profile WHERE email = $1 AND password_hash = $2`
	rows, err := s.pg.Query(ctx, query, username, password)
	if err != nil {
		return nil, fmt.Errorf("pg.Query: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var profileId uuid.UUID
	if err := rows.Scan(&profileId); err != nil {
		return nil, fmt.Errorf("rows.Scan: %w", err)
	}
	return &profileId, nil
}
