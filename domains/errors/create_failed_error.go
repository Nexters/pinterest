package errors

import "fmt"

type CreateFailedError struct {
	Message string
}

func (e CreateFailedError) Error() string {
	return e.Message
}

func NewCreateFailedError(entityName string) CreateFailedError {
	return CreateFailedError{
		Message: fmt.Sprintf("%s %s", entityName, ErrCreationFailed),
	}
}
