package app

type Error struct {
	code     string `json:"code"`
	message  string `json:"message"`
	httpCode int    `json:"http_code"`
}

func BadError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 200,
	}
}

func InternalServerError(code, message string) *Error {
	return &Error{
		code:     code,
		message:  message,
		httpCode: 500,
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
