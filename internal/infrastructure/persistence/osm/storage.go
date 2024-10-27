package osm

import "github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/postgres"

const (
	tableName = "osm_node"
)

type Storage struct {
	pg postgres.PgxPool
}

func New(pg postgres.PgxPool) *Storage {
	return &Storage{
		pg: pg,
	}
}
