// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *BuildRoutesByPointsBadRequest) Validate() error {
	alias := (*Error)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *BuildRoutesByPointsInternalServerError) Validate() error {
	alias := (*Error)(s)
	if err := alias.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *Error) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Status.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ErrorStatus) Validate() error {
	switch s {
	case "not-found":
		return nil
	case "bad-request":
		return nil
	case "internal-error":
		return nil
	case "unauthorized":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
