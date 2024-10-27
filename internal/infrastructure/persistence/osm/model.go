package osm

import "github.com/paulmach/orb"

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
	Coordinates orb.Point
	Type        ObjType
}
