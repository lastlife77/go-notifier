// Package broker provides functionality for working with message brokers.
package broker

import "fmt"

// Error contains the code and the reason.
type Error struct {
	Code   int    // status code
	Reason string // description
}

// NewError creates a new Error with the given code and reason.
func NewError(code int, reason string) *Error {
	return &Error{
		Code:   code,
		Reason: reason,
	}
}

// Error returns the error message.
func (e *Error) Error() string {
	return fmt.Sprintf("%v: %v", e.Code, e.Reason)
}
