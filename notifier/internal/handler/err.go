package handler

import "fmt"

// NotFoundNotifyError represents an error when a requested notify is not found.
type NotFoundNotifyError struct {
	id int
}

// NewNotFoundNotifyError creates a new NotFoundNotifyError with the given event ID.
func NewNotFoundNotifyError(id int) *NotFoundNotifyError {
	return &NotFoundNotifyError{
		id: id,
	}
}

// Error returns the error message for NotFoundNotifyError.
func (e *NotFoundNotifyError) Error() string {
	return fmt.Sprintf("no notify with id found : %v", e.id)
}
