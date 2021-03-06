package app

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestError(t *testing.T) {
	code := fmt.Sprintf("CODE_%d", rand.Int())
	message := fmt.Sprintf("MESSAGE_%d", rand.Int())

	type args struct {
		input    func(code, message string) *Error
		httpCode int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"BadError",
			args{
				input:    BadError,
				httpCode: 200,
			},
		},
		{
			"InternalServerError",
			args{
				input:    InternalServerError,
				httpCode: 500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.args.input(code, message)
			if err.Error() != message {
				t.Error("Message did not match")
			}
			if err.CodeErr() != code {
				t.Error("Code did not match")
			}
			if err.HttpCodeErr() != tt.args.httpCode {
				t.Error("Code did not match")
			}
		})
	}

}
