package route_builder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/paulmach/orb"
)

type RouteBuilder struct {
	httpClient *http.Client
}

func NewRouteBuilder() *RouteBuilder {
	return &RouteBuilder{
		httpClient: http.DefaultClient,
	}
}

func (b *RouteBuilder) BuildRoute(from, to orb.Point, transportType TransportType) ([]Route, error) {
	if err := transportType.Valid(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/route/v1/%s/%f,%f;%f,%f?overview=full&geometries=geojson&annotations=distance&annotations=duration", routerProjectPath, transportType, from.Lon(), from.Lat(), to.Lon(), to.Lat())

	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get: %w", err)
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	response := struct {
		Code   string `json:"code"`
		Routes []struct {
			Geometry struct {
				Coordinates [][]float64 `json:"coordinates"`
				Type        string      `json:"type"`
			} `json:"geometry"`
			Duration float64 `json:"duration"`
			Distance float64 `json:"distance"`
		} `json:"routes"`
	}{}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	if strings.ToLower(response.Code) != "ok" {
		return nil, fmt.Errorf("unexpected code: %s", response.Code)
	}

	routes := make([]Route, 0, len(response.Routes))
	for _, r := range response.Routes {
		routes = append(routes, Route{
			Coordinates: r.Geometry.Coordinates,
			Duration:    r.Duration,
			Distance:    r.Distance,
		})
	}

	return routes, nil
}
