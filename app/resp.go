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

func (r *Resp) MessageResp() string {
	return r.Message
}

func (r *Resp) CodeResp() string {
	return r.Code
}

func (r *Resp) HttpCodeResp() int {
	return r.HttpCode
}
