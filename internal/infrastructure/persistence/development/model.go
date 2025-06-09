package development

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence"
)

type Filter struct {
	SearchQuery    string
	DevelopmentIds *[]int64
	Board          *Board
}

type Pagination struct {
	Page    int
	PerPage int
}

type Board struct {
	BottomRight orb.Point
	TopLeft     orb.Point
}

type Development struct {
	ID          int64
	Name        string
	Coordinates persistence.Point
	CreatedAt   time.Time
	Meta        Meta
}

type Meta struct {
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	AvitoUrl    string `json:"avito_url"`
	GisUrl      string `json:"2gis_url"`
	Address     string `json:"address"`
	Stats       Stats  `json:"stats"`
}

type Stats struct {
	Object3000MCounts POI `json:"object3000mCounts"`
	Distance          POI `json:"distance"`
}

type POI struct {
	Kindergarten int64 `json:"kindergarten"`
	School       int64 `json:"school"`
	Hospital     int64 `json:"hospital"`
	Shops        int64 `json:"shops"`
	Sport        int64 `json:"sport"`
	BusStop      int64 `json:"bus_stop"`
}

func (m *Meta) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, m)
	case string:
		return json.Unmarshal([]byte(v), m)
	default:
		return errors.New("unsupported type for Meta scanning")
	}
}
