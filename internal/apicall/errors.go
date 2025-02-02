package apicall

import (
	"errors"
	"fmt"
)

type APIError struct {
	msg string
}

func (e *APIError) Error() string {
	return e.msg
}
func (e *APIError) Unwrap() error {
	return errors.New(e.msg)
}

// Custom Errors
var ServiceUnavailableErrorType = &APIError{msg: "Service Unavailable"}
var InternalServerErrorType = &APIError{msg: "Internal Server Error"}

/**
 * Throw is a helper function to wrap errors with custom errors
 * @param custom error
 * @param err error
 * @return error
 */
func Throw(custom error, err error) error {
	return fmt.Errorf("%w-%v", custom, err)
}

// Path: internal/apicall/errors.go
