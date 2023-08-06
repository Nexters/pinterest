package errors

import "fmt"

type UpdateFailedError struct {
	Message string
}

func (e UpdateFailedError) Error() string {
	return e.Message
}

func NewUpdateFailedError(entityName string) UpdateFailedError {
	return UpdateFailedError{
		Message: fmt.Sprintf("%s %s", entityName, ErrUpdateFailed),
	}
}
