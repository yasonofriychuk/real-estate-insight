package scanner

import (
	"fmt"
	geojson "github.com/paulmach/go.geojson"
	"github.com/qedus/osmpbf"
	"io"
)

type FeatureScanner struct {
	decoder    *osmpbf.Decoder
	nodeFilter func(*osmpbf.Node) bool
}

func New(osmPbf io.Reader, countGoroutines int) (*FeatureScanner, error) {
	decoder := osmpbf.NewDecoder(osmPbf)
	if err := decoder.Start(countGoroutines); err != nil {
		return nil, fmt.Errorf("could not start osmpbf decoder: %w", err)
	}

	return &FeatureScanner{
		decoder: decoder,
		nodeFilter: func(node *osmpbf.Node) bool {
			v, ok := node.Tags["landuse"]
			return ok && v == "military" //&& node.Tags["name"] != ""
		},
	}, nil
}

func (s *FeatureScanner) Next() (*geojson.Feature, error) {
	for {
		node, err := s.decoder.Decode()
		if err != nil {
			return nil, io.EOF
		}

		n, ok := node.(*osmpbf.Node)
		if !ok || !s.nodeFilter(n) {
			continue
		}

		feature := geojson.NewFeature(geojson.NewPointGeometry([]float64{n.Lon, n.Lat}))
		for k, v := range n.Tags {
			feature.SetProperty(k, v)
		}

		return feature, nil
	}
}
