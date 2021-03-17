package bank

import (
	"context"
	"log"

	"github.com/paytabs/app"
	"github.com/paytabs/bank/delivery"
	bankRepositories "github.com/paytabs/bank/repository"
	"github.com/paytabs/bank/usecase"
	"gopkg.in/paytm/grace.v1"
)

func InitBank(app *app.HttpServer) {

	accounts := make(map[int64]*bankRepositories.AccountDetails)
	accountsRepository := bankRepositories.NewAccountRepository(bankRepositories.AccountsRepositoryParams{AccountsData: accounts})
	transferRepository := bankRepositories.NewTransferRepository(bankRepositories.TransferRepositoryParams{AccountsData: accounts})
	dataRepository := bankRepositories.NewDataRepository(bankRepositories.DataRepositoryParams{AccountsData: accounts})

	u := usecase.NewUsecase(usecase.UseCaseInit{
		TransferRepository: transferRepository,
		AccountsRepository: accountsRepository,
		DataRepository:     dataRepository,
	})

	delivery.NewHTTPDelivery(u, app)

	log.Println(context.Background(), grace.Serve("localhost:8080", app))

}
