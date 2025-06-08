package location_list

import (
	"context"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/persistence/location"
)

type storage interface {
	CityList(ctx context.Context) ([]location.Location, error)
}
