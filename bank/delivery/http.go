package delivery

import (
	"github.com/paytabs/app"
	"github.com/paytabs/bank/usecase"
)

type bankTransferDelivery struct {
	useCase    usecase.UseCase
	httpServer *app.HttpServer
}

func NewHTTPDelivery(useCase usecase.UseCase, httpServer *app.HttpServer) {
	p := &bankTransferDelivery{
		useCase:    useCase,
		httpServer: httpServer,
	}

	InitHTTPEndpoint(p)
}

func InitHTTPEndpoint(delivery *bankTransferDelivery) {
}
