package development

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/paulmach/orb"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence"
)

type Filter struct {
	SearchQuery string
	Board       *Board
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
