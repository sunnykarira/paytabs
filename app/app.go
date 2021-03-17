package app

import (
	"os"
)

type App struct {
	Hostname  string
	externals map[string]interface{}
}

func NewApp() App {
	host, err := os.Hostname()
	if err != nil {
		host = "undefined"
	}

	return App{
		Hostname: host,
	}
}

func (a *App) SetData(key string, value interface{}) {
	a.externals = make(map[string]interface{})
	a.externals[key] = value
}

func (a *App) GetData(key string) interface{} {

	val, ok := a.externals[key]
	if !ok {
		return nil
	}
	return val
}
