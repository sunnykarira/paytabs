package bank

import (
	"context"

	"github.com/paytabs/app"
	bankRepositories "github.com/paytabs/bank/repository"
	"github.com/paytabs/bank/usecase"
)

func InitBank(app *app.HttpServer) {

	ctx := context.Background()

	accounts := make(map[int64]bankRepositories.AccountDetails)
	accountsRepository := bankRepositories.NewAccountRepository(bankRepositories.AccountsRepositoryParams{AccountsData: accounts})
	transferRepository := bankRepositories.NewTransferRepository(bankRepositories.TransferRepositoryParams{})

	u := usecase.NewUsecase(usecase.UseCaseInit{
		TransferRepository: transferRepository,
		AccountsRepository: accountsRepository,
	})

	delivery.NewDelivery(u, app)

}
