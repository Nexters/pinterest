package errors

import "fmt"

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(entityName string) NotFoundError {
	return NotFoundError{
		Message: fmt.Sprintf("%s %s", entityName, ErrNotFound),
	}
}
