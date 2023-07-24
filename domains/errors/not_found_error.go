package errors

import "fmt"

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(entityName string) error {
	return &NotFoundError{
		Message: fmt.Sprintf("해당하는 %s를 찾을 수 없습니다.", entityName),
	}
}
