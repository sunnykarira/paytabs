package usecase

import (
	"context"
	"github.com/paytabs/bank/model"
	"time"
)

func (u *usecase) FetchData(ctx context.Context, accountID int64) (accountDetails model.Account, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return u.accountsRepo.FetchData(newCtx, accountID)
}

func (u *usecase) CreateAccount(ctx context.Context, accountData model.Account) (success bool, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return u.accountsRepo.CreateAccount(newCtx, accountData)
}

func (u *usecase) AddMoney(ctx context.Context, accountID int64, money float64) (success bool, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return u.accountsRepo.AddMoney(newCtx, accountID, money)
}

func (u *usecase) SendMoney(ctx context.Context, txnData model.TransactionData) (success bool, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return u.transferRepo.SendMoney(newCtx, txnData)
}
