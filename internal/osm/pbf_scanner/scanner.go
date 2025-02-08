package pbf_scanner

import (
	"fmt"
	"io"

	"github.com/qedus/osmpbf"
)

type FeatureScanner struct {
	decoder *osmpbf.Decoder
}

func New(osmPbf io.Reader, countGoroutines int) (*FeatureScanner, error) {
	decoder := osmpbf.NewDecoder(osmPbf)
	if err := decoder.Start(countGoroutines); err != nil {
		return nil, fmt.Errorf("could not start osmpbf decoder: %w", err)
	}

	return &FeatureScanner{
		decoder: decoder,
	}, nil
}

func (s *FeatureScanner) Next() (osmpbf.Node, error) {
	for {
		node, err := s.decoder.Decode()
		if err != nil {
			return osmpbf.Node{}, io.EOF
		}

		n, ok := node.(*osmpbf.Node)
		if !ok {
			continue
		}

		return *n, nil
	}
}
