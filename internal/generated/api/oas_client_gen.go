// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// BuildRoutesByPoints invokes buildRoutesByPoints operation.
	//
	// Build a route between points.
	//
	// GET /routes/build/points
	BuildRoutesByPoints(ctx context.Context, params BuildRoutesByPointsParams) (BuildRoutesByPointsRes, error)
	// ObjectsFindNearestInfrastructure invokes objectsFindNearestInfrastructure operation.
	//
	// Search for nearby infrastructure facilities.
	//
	// GET /objects/find/nearestInfrastructure
	ObjectsFindNearestInfrastructure(ctx context.Context, params ObjectsFindNearestInfrastructureParams) (ObjectsFindNearestInfrastructureRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// BuildRoutesByPoints invokes buildRoutesByPoints operation.
//
// Build a route between points.
//
// GET /routes/build/points
func (c *Client) BuildRoutesByPoints(ctx context.Context, params BuildRoutesByPointsParams) (BuildRoutesByPointsRes, error) {
	res, err := c.sendBuildRoutesByPoints(ctx, params)
	return res, err
}

func (c *Client) sendBuildRoutesByPoints(ctx context.Context, params BuildRoutesByPointsParams) (res BuildRoutesByPointsRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("buildRoutesByPoints"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/routes/build/points"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, BuildRoutesByPointsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/routes/build/points"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "latFrom" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "latFrom",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.LatFrom.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "lonFrom" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "lonFrom",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.LonFrom.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "latTo" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "latTo",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.LatTo.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "lonTo" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "lonTo",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.LonTo.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeBuildRoutesByPointsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ObjectsFindNearestInfrastructure invokes objectsFindNearestInfrastructure operation.
//
// Search for nearby infrastructure facilities.
//
// GET /objects/find/nearestInfrastructure
func (c *Client) ObjectsFindNearestInfrastructure(ctx context.Context, params ObjectsFindNearestInfrastructureParams) (ObjectsFindNearestInfrastructureRes, error) {
	res, err := c.sendObjectsFindNearestInfrastructure(ctx, params)
	return res, err
}

func (c *Client) sendObjectsFindNearestInfrastructure(ctx context.Context, params ObjectsFindNearestInfrastructureParams) (res ObjectsFindNearestInfrastructureRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("objectsFindNearestInfrastructure"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/objects/find/nearestInfrastructure"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ObjectsFindNearestInfrastructureOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/objects/find/nearestInfrastructure"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "lat" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "lat",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Lat.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "lon" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "lon",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Lon.Get(); ok {
				return e.EncodeValue(conv.Float64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "objectTypes" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "objectTypes",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.ObjectTypes != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.ObjectTypes {
						if err := func() error {
							return e.EncodeValue(conv.StringToString(string(item)))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeObjectsFindNearestInfrastructureResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
