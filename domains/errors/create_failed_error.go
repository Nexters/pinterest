package errors

import "fmt"

type CreateFailedError struct {
	Message string
}

func (e *CreateFailedError) Error() string {
	return e.Message
}

func NewCreateFailedError(entityName string) error {
	return &CreateFailedError{
		Message: fmt.Sprintf("%s 생성에 실패하였습니다.", entityName),
	}
}
