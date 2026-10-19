package httperr

import (
	"fmt"
	"net/http"
	"slices"
)

type HttpError struct {
	Cause      string       `json:"cause"`
	StatusCode int          `json:"-"`
	Fields     []FieldError `json:"fields,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Cause
}

func NewHttpError(cause string, statusCode int) *HttpError {
	return &HttpError{
		Cause:      cause,
		StatusCode: statusCode,
	}
}

func (e *HttpError) AppendField(field, cause string) {
	e.Fields = append(e.Fields, FieldError{Field: field, Cause: cause})
}

func (e *HttpError) HasFields() bool {
	return len(e.Fields) > 0
}

func ErrBadRequest() *HttpError {
	return NewHttpError("can't process request", http.StatusBadRequest)
}

func ErrUnauthorized() *HttpError {
	return NewHttpError("you're not authenticated", http.StatusUnauthorized)
}

func ErrForbidden() *HttpError {
	return NewHttpError("can't access this resource", http.StatusForbidden)
}

func ErrNotFound(resource string) *HttpError {
	return NewHttpError(fmt.Sprintf("%s not found or not available", resource), http.StatusNotFound)
}

func ErrConflict() *HttpError {
	return NewHttpError("resource conflict", http.StatusConflict)
}

func ErrUnprocessableEntity(fields ...FieldError) *HttpError {
	e := NewHttpError("entity validation error", http.StatusUnprocessableEntity)
	e.Fields = slices.Clone(fields)
	return e
}

func ErrInternalServerError() *HttpError {
	return NewHttpError("server internal error occurred", http.StatusInternalServerError)
}
