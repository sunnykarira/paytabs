package bank

import (
	"github.com/paytabs/app"
	bankRepositories "github.com/paytabs/bank/repository"
	"github.com/paytabs/bank/usecase"
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

	delivery.NewDelivery(u, app)

}
