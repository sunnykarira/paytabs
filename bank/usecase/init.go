package usecase

import (
	"github.com/paytabs/bank/repository"
)

type usecase struct {
	transferRepo repository.TransferRepository
	accountsRepo repository.AccountsRepository
	dataRepo     repository.DataRepository
}

func NewUsecase(u UseCaseInit) UseCase {
	return &usecase{
		transferRepo: u.TransferRepository,
		accountsRepo: u.AccountsRepository,
		dataRepo:     u.DataRepository,
	}
}
