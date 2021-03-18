package app

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestResp(t *testing.T) {
	code := fmt.Sprintf("CODE_%d", rand.Int())
	message := fmt.Sprintf("MESSAGE_%d", rand.Int())

	type args struct {
		input    func(code, message string) *Resp
		httpCode int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"Success",
			args{
				input:    RespMessage,
				httpCode: 200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.input(code, message)
			if err.MessageResp() != message {
				t.Error("Message did not match")
			}
			if err.CodeResp() != code {
				t.Error("Code did not match")
			}
		})
	}

}
