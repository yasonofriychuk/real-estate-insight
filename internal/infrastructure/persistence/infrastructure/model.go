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

type HeatmapParams struct {
	BBox     HeatmapParamsBBox    `json:"bbox"`
	Weights  HeatmapParamsWeights `json:"weights"`
	CellSize int                  `json:"cellSize" db:"cell_size_meters"`
}

type HeatmapParamsBBox struct {
	TopLeftLon     float64 `json:"topLeftLon" db:"top_left_lon"`
	TopLeftLat     float64 `json:"topLeftLat" db:"top_left_lat"`
	BottomRightLon float64 `json:"bottomRightLon" db:"bottom_right_lon"`
	BottomRightLat float64 `json:"bottomRightLat" db:"bottom_right_lat"`
}

type HeatmapParamsWeights struct {
	Hospital     int `json:"hospital" db:"w_hospital"`
	Sport        int `json:"sport" db:"w_sport"`
	Shops        int `json:"shops" db:"w_shops"`
	Kindergarten int `json:"kindergarten" db:"w_kindergarten"`
	BusStop      int `json:"bus_stop" db:"w_bus_stop"`
	School       int `json:"school" db:"w_school"`
}

type HexagonWeight struct {
	Geometry    map[string]interface{} `json:"geometry"`     // GeoJSON Polygon
	TotalWeight float64                `json:"total_weight"` // агрегированный вес
}
