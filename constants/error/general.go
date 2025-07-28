package error

import "errors"

var (
	ErrInternalServer      = errors.New("internal server error")
	ErrTooManyRequest      = errors.New("too many request")
	ErrBadGateway          = errors.New("bad gateway")
	ErrServiceUnavailable  = errors.New("service unavailable")
	ErrGatewayTimeout      = errors.New("gateway timeout")
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrInvalidToken        = errors.New("invalid token")
	ErrSQLError            = errors.New("sql error")
)

var GeneralErrors = []error{
	ErrInternalServer,
	ErrTooManyRequest,
	ErrBadGateway,
	ErrServiceUnavailable,
	ErrGatewayTimeout,
	ErrInternalServerError,
	ErrNotFound,
	ErrBadRequest,
	ErrUnauthorized,
	ErrForbidden,
	ErrInvalidToken,
}
