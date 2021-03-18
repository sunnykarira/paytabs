package app

type Resp struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"http_code"`
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
