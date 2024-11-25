// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 */

package openapi

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrTypeAssertionError is thrown when type an interface does not match the asserted type
	ErrTypeAssertionError = errors.New("unable to assert type")
)

// ParsingError indicates that an error has occurred when parsing request parameters
type ParsingError struct {
	Param string
	Err   error
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}

func (e *ParsingError) Error() string {
	if e.Param == "" {
		return e.Err.Error()
	}

	return e.Param + ": " + e.Err.Error()
}

// RequiredError indicates that an error has occurred when parsing request parameters
type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("required field '%s' is zero value.", e.Field)
}

// ErrorHandler defines the required method for handling error. You may implement it and inject this into a controller if
// you would like errors to be handled differently from the DefaultErrorHandler
type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error, result *ImplResponse)

// DefaultErrorHandler defines the default logic on how to handle errors from the controller. Any errors from parsing
// request params will return a StatusBadRequest. Otherwise, the error code originating from the servicer will be used.
func DefaultErrorHandler(w http.ResponseWriter, _ *http.Request, err error, result *ImplResponse) {
	var parsingErr *ParsingError
	if ok := errors.As(err, &parsingErr); ok {
		// Handle parsing errors
		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusBadRequest), w)
		return
	}

	var requiredErr *RequiredError
	if ok := errors.As(err, &requiredErr); ok {
		// Handle missing required errors
		_ = EncodeJSONResponse(err.Error(), func(i int) *int { return &i }(http.StatusUnprocessableEntity), w)
		return
	} 

	// Handle all other errors
	_ = EncodeJSONResponse(err.Error(), &result.Code, w)
}
