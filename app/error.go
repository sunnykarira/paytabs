package app

type Error struct {
	code     string
	message  string
	httpCode int
}

func BadError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 400,
	}
}

func NotFoundError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 404,
	}
}

func InternalServerError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 500,
	}
}

func ForbiddenError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 403,
	}
}

func (err *Error) Error() string {
	return err.message
}

func (err *Error) Code() string {
	return err.code
}

func (err *Error) HttpCode() int {
	return err.httpCode
}
