package app

type Error struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"http_code"`
}

func BadError(code, message string) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		HttpCode: 200,
	}
}

func InternalServerError(code, message string) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		HttpCode: 500,
	}
}

func (err *Error) Error() string {
	return err.Message
}

func (err *Error) CodeErr() string {
	return err.Code
}

func (err *Error) HttpCodeErr() int {
	return err.HttpCode
}
