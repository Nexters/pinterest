package errors

type UnauthorizedError struct {
	Message string
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

func NewUnauthorizedError() UnauthorizedError {
	return UnauthorizedError{
		Message: ErrUnauthorized,
	}
}
