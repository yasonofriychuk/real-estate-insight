package infrastructure

import (
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence"
	"github.com/yasonofriychuk/real-estate-insight/internal/osm"
)

type Obj struct {
	ID          int
	Name        string
	Coordinates persistence.Point
	Type        osm.ObjType
}
