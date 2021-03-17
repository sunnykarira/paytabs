package repository

import (
	"context"
	"github.com/paytabs/cmd/model"
	"sync"
)

type (
	AccountDetails struct {
		model.Account
		mu sync.Mutex
	}

	AccountsRepositoryParams struct {
		dataStore map[int64]AccountDetails
	}

	accountsRepo struct {
		accountsData map[int64]AccountDetails
	}
)


func NewAccountRepository(params AccountsRepositoryParams) AccountsRepository {

	return &accountsRepo{
		accountsData: params.dataStore,
	}
}

func (a *accountsRepo) FetchData(ctx context.Context, accountID int64) (accountDetails model.Account, err error){

	return model.Account{}, nil
}

func (a *accountsRepo) CreateAccount(ctx context.Context, accountData model.Account) (success bool, err error){

	return false, nil
}
func (a *accountsRepo) AddMoney(ctx context.Context, accountID int64, money float64) (success bool, err error){

	return  false, nil
}
