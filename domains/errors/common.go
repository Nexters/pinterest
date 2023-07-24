package errors

import "errors"

const (
	ErrNotFound     = "해당하는 값을 찾을 수 없습니다"
	ErrCreateFailed = "생성에 실패했습니다"
)

func NewNotFoundError() error {
	return errors.New(ErrNotFound)
}

func NewCreateFailedError() error {
	return errors.New(ErrCreateFailed)
}
