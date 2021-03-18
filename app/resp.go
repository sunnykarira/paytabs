package app

type Resp struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
}

func RespMessage(code, message string) *Resp {
	return &Resp{
		Code:     code,
		Message:  message,
	}
}

func (r *Resp) MessageResp() string {
	return r.Message
}

func (r *Resp) CodeResp() string {
	return r.Code
}

