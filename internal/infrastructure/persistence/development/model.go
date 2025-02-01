package development

import (
	"time"

	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence"
)

type Development struct {
	ID          int
	Name        string
	Coordinates persistence.Point
	CreatedAt   time.Time
}
