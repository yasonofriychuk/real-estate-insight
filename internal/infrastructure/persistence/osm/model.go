package osm

import (
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence"
)

type ObjType string

const (
	Hospital     ObjType = "Hospital"
	Sport        ObjType = "Sport"
	Shops        ObjType = "Shops"
	Kindergarten ObjType = "Kindergarten"
	BusStop      ObjType = "BusStop"
)

type Obj struct {
	Name        string
	Coordinates persistence.Point
	Type        ObjType
}
