package persistence

import (
	"errors"
	"fmt"

	geojson "github.com/paulmach/go.geojson"
	"github.com/paulmach/orb"
)

type Point struct {
	orb.Point
}

func (p *Point) Scan(value any) error {
	source, ok := value.(string)
	if !ok {
		return errors.New("incompatible type")
	}

	geometry, err := geojson.UnmarshalGeometry([]byte(source))
	if err != nil {
		return fmt.Errorf("unmarshal geometry: %w", err)
	}

	if geometry.Type != geojson.GeometryPoint {
		return fmt.Errorf("invalid geometry type: %s", geometry.Type)
	}

	p.Point = orb.Point(geometry.Point)
	return nil
}
