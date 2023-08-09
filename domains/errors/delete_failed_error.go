package errors

import "fmt"

type DeleteFailedError struct {
	Message string
}

func (e DeleteFailedError) Error() string {
	return e.Message
}

func NewDeleteFailedError(entityName string) DeleteFailedError {
	return DeleteFailedError{
		Message: fmt.Sprintf("%s %s", entityName, ErrDeletionFailed),
	}
}
