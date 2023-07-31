package errors

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func NewUnauthorizedError() error {
	return &UnauthorizedError{
		Message: "ID 혹은 비밀번호가 잘못되었습니다.",
	}
}
