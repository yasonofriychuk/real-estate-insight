package route_builder

import (
	"fmt"
	"slices"
)

const (
	routerProjectPath = "http://router.project-osrm.org"
)

type TransportType string

const (
	CarTransportType  TransportType = "car"
	BikeTransportType TransportType = "bike"
	FootTransportType TransportType = "foot"
)

var (
	TransportTypes = []TransportType{CarTransportType, BikeTransportType, FootTransportType}
)

func (t TransportType) Valid() error {
	if !slices.Contains([]TransportType{CarTransportType, BikeTransportType, FootTransportType}, t) {
		return fmt.Errorf("invalid TransportType: %s", t)
	}
	return nil
}

type Route struct {
	Coordinates [][]float64
	Duration    float64
	Distance    float64
}
