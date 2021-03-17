package main

import (
	"context"
	"github.com/paytabs/app"
)

func main() {

	flgs := ParseFlags()

	ctx := context.Background()
	application := app.NewApp()

	server := app.NewHttpServer(application)

	InitServices(ctx, server)

}


func InitServices(ctx context.Context, server *app.HttpServer) {


}