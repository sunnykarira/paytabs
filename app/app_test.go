package app

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	host, err := os.Hostname()
	if err != nil {
		host = "undefined"
	}

	app := NewApp()
	assert.Equal(t, app.Hostname, host)
}

func TestApp_SetData(t *testing.T) {
	host, err := os.Hostname()
	if err != nil {
		host = "undefined"
	}

	type fields struct {
		Hostname  string
		externals map[string]interface{}
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "when params are fine",
			fields: fields{Hostname: host},
			args:   args{key: "key", value: "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Hostname:  tt.fields.Hostname,
				externals: tt.fields.externals,
			}
			a.SetData(tt.args.key, tt.args.value)
		})
	}
}

func TestApp_GetData(t *testing.T) {
	host, err := os.Hostname()
	if err != nil {
		host = "undefined"
	}

	mockMap := make(map[string]interface{})
	mockMap["key"] = "value"

	type fields struct {
		Hostname  string
		externals map[string]interface{}
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name:   "When key exists",
			fields: fields{Hostname: host, externals: mockMap},
			args:   args{key: "key"},
			want:   "value",
		},
		{
			name:   "When key not exists",
			fields: fields{Hostname: host, externals: mockMap},
			args:   args{key: "key1"},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{
				Hostname:  tt.fields.Hostname,
				externals: tt.fields.externals,
			}
			if got := a.GetData(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
