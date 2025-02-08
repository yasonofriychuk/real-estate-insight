package errors

import "github.com/yasonofriychuk/real-estate-insight/internal/generated/api"

func BuildError(code int, message string) api.Error {
	var status api.ErrorStatus

	switch code {
	case 400:
		status = api.ErrorStatusBadRequest
	case 403:
		status = api.ErrorStatusUnauthorized
	case 404:
		status = api.ErrorStatusNotFound
	case 500:
		status = api.ErrorStatusInternalError
	default:
		status = api.ErrorStatusBadRequest
	}

	return api.Error{
		Status: status,
		Error: api.ErrorError{
			Code:    code,
			Message: message,
		},
	}
}
