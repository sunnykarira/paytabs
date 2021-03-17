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