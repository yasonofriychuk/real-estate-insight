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

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// AddToFavoriteSelection invokes addToFavoriteSelection operation.
	//
	// Add or remove a development to/from the selected user's favorite selection.
	//
	// POST /selection/favorite
	AddToFavoriteSelection(ctx context.Context, request *AddToFavoriteSelectionReq) (AddToFavoriteSelectionRes, error)
	// BuildRoutesByPoints invokes buildRoutesByPoints operation.
	//
	// Build a route between points.
	//
	// GET /routes/build/points
	BuildRoutesByPoints(ctx context.Context, params BuildRoutesByPointsParams) (BuildRoutesByPointsRes, error)
	// CreateSelection invokes createSelection operation.
	//
	// Create a new selection for the user with name, comment, and form.
	//
	// POST /selection/create
	CreateSelection(ctx context.Context, request *CreateSelectionReq) (CreateSelectionRes, error)
	// DeleteSelection invokes deleteSelection operation.
	//
	// Delete a selection for the user by selection ID.
	//
	// POST /selection/delete
	DeleteSelection(ctx context.Context, params DeleteSelectionParams) (DeleteSelectionRes, error)
	// DevelopmentSearch invokes developmentSearch operation.
	//
	// POST /developments/search
	DevelopmentSearch(ctx context.Context, request *DevelopmentSearchReq) (DevelopmentSearchRes, error)
	// EditSelection invokes editSelection operation.
	//
	// Edit new selection.
	//
	// POST /selection/edit
	EditSelection(ctx context.Context, request *EditSelectionReq) (EditSelectionRes, error)
	// GenerateInfrastructureHeatmap invokes generateInfrastructureHeatmap operation.
	//
	// Returns a grid-based heatmap for infrastructure objects based on type weights within a selected
	// bounding box.
	//
	// POST /infrastructure/heatmap
	GenerateInfrastructureHeatmap(ctx context.Context, request *GenerateInfrastructureHeatmapReq) (GenerateInfrastructureHeatmapRes, error)
	// InfrastructureRadiusBoard invokes infrastructureRadiusBoard operation.
	//
	// Search for infrastructure around the selected residential complex.
	//
	// GET /infrastructure/radius
	InfrastructureRadiusBoard(ctx context.Context, params InfrastructureRadiusBoardParams) (InfrastructureRadiusBoardRes, error)
	// LocationList invokes locationList operation.
	//
	// Get location list.
	//
	// GET /location/list
	LocationList(ctx context.Context) (LocationListRes, error)
	// SelectionById invokes selectionById operation.
	//
	// Get selection.
	//
	// GET /selection/{selectionId}
	SelectionById(ctx context.Context, params SelectionByIdParams) (SelectionByIdRes, error)
	// SelectionList invokes selectionList operation.
	//
	// Get selection list.
	//
	// GET /selection/list
	SelectionList(ctx context.Context) (SelectionListRes, error)
	// UserLogin invokes userLogin operation.
	//
	// Authenticate the user using email and password.
	//
	// POST /profile/login
	UserLogin(ctx context.Context, request *UserLoginReq) (UserLoginRes, error)
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

var _ Handler = struct {
	*Client
}{}

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

// AddToFavoriteSelection invokes addToFavoriteSelection operation.
//
// Add or remove a development to/from the selected user's favorite selection.
//
// POST /selection/favorite
func (c *Client) AddToFavoriteSelection(ctx context.Context, request *AddToFavoriteSelectionReq) (AddToFavoriteSelectionRes, error) {
	res, err := c.sendAddToFavoriteSelection(ctx, request)
	return res, err
}

func (c *Client) sendAddToFavoriteSelection(ctx context.Context, request *AddToFavoriteSelectionReq) (res AddToFavoriteSelectionRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("addToFavoriteSelection"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/selection/favorite"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, AddToFavoriteSelectionOperation,
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
	pathParts[0] = "/selection/favorite"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAddToFavoriteSelectionRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAddToFavoriteSelectionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
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
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
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
		// Encode "developmentId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "developmentId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.Int64ToString(params.DevelopmentId))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "osmId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "osmId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.Int64ToString(params.OsmId))
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

// CreateSelection invokes createSelection operation.
//
// Create a new selection for the user with name, comment, and form.
//
// POST /selection/create
func (c *Client) CreateSelection(ctx context.Context, request *CreateSelectionReq) (CreateSelectionRes, error) {
	res, err := c.sendCreateSelection(ctx, request)
	return res, err
}

func (c *Client) sendCreateSelection(ctx context.Context, request *CreateSelectionReq) (res CreateSelectionRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSelection"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/selection/create"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateSelectionOperation,
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
	pathParts[0] = "/selection/create"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateSelectionRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateSelectionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteSelection invokes deleteSelection operation.
//
// Delete a selection for the user by selection ID.
//
// POST /selection/delete
func (c *Client) DeleteSelection(ctx context.Context, params DeleteSelectionParams) (DeleteSelectionRes, error) {
	res, err := c.sendDeleteSelection(ctx, params)
	return res, err
}

func (c *Client) sendDeleteSelection(ctx context.Context, params DeleteSelectionParams) (res DeleteSelectionRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteSelection"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/selection/delete"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteSelectionOperation,
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
	pathParts[0] = "/selection/delete"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.UUIDToString(params.ID))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
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
	result, err := decodeDeleteSelectionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DevelopmentSearch invokes developmentSearch operation.
//
// POST /developments/search
func (c *Client) DevelopmentSearch(ctx context.Context, request *DevelopmentSearchReq) (DevelopmentSearchRes, error) {
	res, err := c.sendDevelopmentSearch(ctx, request)
	return res, err
}

func (c *Client) sendDevelopmentSearch(ctx context.Context, request *DevelopmentSearchReq) (res DevelopmentSearchRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("developmentSearch"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/developments/search"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DevelopmentSearchOperation,
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
	pathParts[0] = "/developments/search"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeDevelopmentSearchRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeDevelopmentSearchResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// EditSelection invokes editSelection operation.
//
// Edit new selection.
//
// POST /selection/edit
func (c *Client) EditSelection(ctx context.Context, request *EditSelectionReq) (EditSelectionRes, error) {
	res, err := c.sendEditSelection(ctx, request)
	return res, err
}

func (c *Client) sendEditSelection(ctx context.Context, request *EditSelectionReq) (res EditSelectionRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("editSelection"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/selection/edit"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, EditSelectionOperation,
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
	pathParts[0] = "/selection/edit"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeEditSelectionRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeEditSelectionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GenerateInfrastructureHeatmap invokes generateInfrastructureHeatmap operation.
//
// Returns a grid-based heatmap for infrastructure objects based on type weights within a selected
// bounding box.
//
// POST /infrastructure/heatmap
func (c *Client) GenerateInfrastructureHeatmap(ctx context.Context, request *GenerateInfrastructureHeatmapReq) (GenerateInfrastructureHeatmapRes, error) {
	res, err := c.sendGenerateInfrastructureHeatmap(ctx, request)
	return res, err
}

func (c *Client) sendGenerateInfrastructureHeatmap(ctx context.Context, request *GenerateInfrastructureHeatmapReq) (res GenerateInfrastructureHeatmapRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("generateInfrastructureHeatmap"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/infrastructure/heatmap"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GenerateInfrastructureHeatmapOperation,
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
	pathParts[0] = "/infrastructure/heatmap"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeGenerateInfrastructureHeatmapRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeGenerateInfrastructureHeatmapResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InfrastructureRadiusBoard invokes infrastructureRadiusBoard operation.
//
// Search for infrastructure around the selected residential complex.
//
// GET /infrastructure/radius
func (c *Client) InfrastructureRadiusBoard(ctx context.Context, params InfrastructureRadiusBoardParams) (InfrastructureRadiusBoardRes, error) {
	res, err := c.sendInfrastructureRadiusBoard(ctx, params)
	return res, err
}

func (c *Client) sendInfrastructureRadiusBoard(ctx context.Context, params InfrastructureRadiusBoardParams) (res InfrastructureRadiusBoardRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("infrastructureRadiusBoard"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/infrastructure/radius"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, InfrastructureRadiusBoardOperation,
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
	pathParts[0] = "/infrastructure/radius"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "developmentId" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "developmentId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.DevelopmentId))
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "radius" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "radius",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			return e.EncodeValue(conv.IntToString(params.Radius))
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
	result, err := decodeInfrastructureRadiusBoardResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LocationList invokes locationList operation.
//
// Get location list.
//
// GET /location/list
func (c *Client) LocationList(ctx context.Context) (LocationListRes, error) {
	res, err := c.sendLocationList(ctx)
	return res, err
}

func (c *Client) sendLocationList(ctx context.Context) (res LocationListRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("locationList"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/location/list"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, LocationListOperation,
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
	pathParts[0] = "/location/list"
	uri.AddPathParts(u, pathParts[:]...)

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
	result, err := decodeLocationListResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SelectionById invokes selectionById operation.
//
// Get selection.
//
// GET /selection/{selectionId}
func (c *Client) SelectionById(ctx context.Context, params SelectionByIdParams) (SelectionByIdRes, error) {
	res, err := c.sendSelectionById(ctx, params)
	return res, err
}

func (c *Client) sendSelectionById(ctx context.Context, params SelectionByIdParams) (res SelectionByIdRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("selectionById"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/selection/{selectionId}"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SelectionByIdOperation,
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
	var pathParts [2]string
	pathParts[0] = "/selection/"
	{
		// Encode "selectionId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "selectionId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.UUIDToString(params.SelectionId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

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
	result, err := decodeSelectionByIdResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SelectionList invokes selectionList operation.
//
// Get selection list.
//
// GET /selection/list
func (c *Client) SelectionList(ctx context.Context) (SelectionListRes, error) {
	res, err := c.sendSelectionList(ctx)
	return res, err
}

func (c *Client) sendSelectionList(ctx context.Context) (res SelectionListRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("selectionList"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/selection/list"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SelectionListOperation,
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
	pathParts[0] = "/selection/list"
	uri.AddPathParts(u, pathParts[:]...)

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
	result, err := decodeSelectionListResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UserLogin invokes userLogin operation.
//
// Authenticate the user using email and password.
//
// POST /profile/login
func (c *Client) UserLogin(ctx context.Context, request *UserLoginReq) (UserLoginRes, error) {
	res, err := c.sendUserLogin(ctx, request)
	return res, err
}

func (c *Client) sendUserLogin(ctx context.Context, request *UserLoginReq) (res UserLoginRes, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("userLogin"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/profile/login"),
	}

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UserLoginOperation,
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
	pathParts[0] = "/profile/login"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUserLoginRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUserLoginResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
