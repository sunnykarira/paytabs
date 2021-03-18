package main

import (
	"github.com/paytabs/app"
	"github.com/paytabs/bank"
	"time"
)

func main() {

	application := app.NewApp()

	server := app.NewHttpServer(application)

	go func() {
		time.Sleep(3 * time.Second)
		RunCURLScript()
	}()

	InitService(server)

}

func InitService(server *app.HttpServer) {

	bank.InitBank(server)

}
