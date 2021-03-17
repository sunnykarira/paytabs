package main

import (
	"github.com/paytabs/app"
	"github.com/paytabs/bank"
)

func main() {

	application := app.NewApp()

	server := app.NewHttpServer(application)

	InitService(server)

}

func InitService(server *app.HttpServer) {

	bank.InitBank(server)

}
