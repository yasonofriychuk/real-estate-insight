// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// BuildRoutesByPointsParams is parameters of buildRoutesByPoints operation.
type BuildRoutesByPointsParams struct {
	DevelopmentId int64
	OsmId         int64
}

func unpackBuildRoutesByPointsParams(packed middleware.Parameters) (params BuildRoutesByPointsParams) {
	{
		key := middleware.ParameterKey{
			Name: "developmentId",
			In:   "query",
		}
		params.DevelopmentId = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "osmId",
			In:   "query",
		}
		params.OsmId = packed[key].(int64)
	}
	return params
}

func decodeBuildRoutesByPointsParams(args [0]string, argsEscaped bool, r *http.Request) (params BuildRoutesByPointsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: developmentId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "developmentId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.DevelopmentId = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "developmentId",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: osmId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "osmId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.OsmId = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "osmId",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// InfrastructureRadiusBoardParams is parameters of infrastructureRadiusBoard operation.
type InfrastructureRadiusBoardParams struct {
	DevelopmentId int
	Radius        int
}

func unpackInfrastructureRadiusBoardParams(packed middleware.Parameters) (params InfrastructureRadiusBoardParams) {
	{
		key := middleware.ParameterKey{
			Name: "developmentId",
			In:   "query",
		}
		params.DevelopmentId = packed[key].(int)
	}
	{
		key := middleware.ParameterKey{
			Name: "radius",
			In:   "query",
		}
		params.Radius = packed[key].(int)
	}
	return params
}

func decodeInfrastructureRadiusBoardParams(args [0]string, argsEscaped bool, r *http.Request) (params InfrastructureRadiusBoardParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: developmentId.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "developmentId",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.DevelopmentId = c
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           0,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.DevelopmentId)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "developmentId",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: radius.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "radius",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt(val)
				if err != nil {
					return err
				}

				params.Radius = c
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1000,
					MaxSet:        true,
					Max:           10000,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.Radius)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "radius",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
