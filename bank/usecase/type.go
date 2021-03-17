package usecase

import "github.com/paytabs/bank/repository"

type UseCaseInit struct{
	TransferRepository repository.TransferRepository
	AccountsRepository repository.AccountsRepository
}