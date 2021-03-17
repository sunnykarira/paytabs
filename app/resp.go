package app

type Resp struct {
	Code     string
	Message  string
	HttpCode int
}

func RespMessage(code, message string) *Resp {
	return &Resp{
		Code:     code,
		Message:  message,
		HttpCode: 200,
	}
}
