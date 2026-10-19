package httperr

type FieldError struct {
	Field string `json:"field"`
	Cause string `json:"cause"`
}

func NewFieldError(field, cause string) *FieldError {
	return &FieldError{
		Field: field,
		Cause: cause,
	}
}
