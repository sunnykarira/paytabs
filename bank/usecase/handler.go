package usecase

import (
	"context"
	"errors"
	"github.com/paytabs/bank/model"
	"strconv"
	"sync"
	"time"
)

func (u *usecase) FetchData(ctx context.Context, accountID int64) (accountDetails model.Account, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return u.dataRepo.FetchData(newCtx, accountID)
}

func (u *usecase) CreateAccount(ctx context.Context, accountData model.Account) (success bool, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return u.accountsRepo.CreateAccount(newCtx, accountData)
}

func (u *usecase) AddMoney(ctx context.Context, accountID int64, money float64) (success bool, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	return u.accountsRepo.AddMoney(newCtx, accountID, money)
}

func (u *usecase) SendMoney(ctx context.Context, txnData model.TransactionData) (success bool, err error) {

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var (
		srcData  model.Account
		destData model.Account
	)
	var (
		srcErr  error
		destErr error
	)

	var wg *sync.WaitGroup
	wg.Add(2)
	go func() {
		srcData, srcErr = u.dataRepo.FetchData(newCtx, txnData.SourceAccountID)
	}()

	go func() {
		destData, destErr = u.dataRepo.FetchData(newCtx, txnData.DestinationAccountID)
	}()
	wg.Wait()

	if srcErr != nil {
		return false, srcErr
	}
	if destErr != nil {
		return false, destErr
	}
	if srcData.Balance-txnData.Amount < 0 {
		return false, errors.New("not enough balance in account to send money from " + strconv.FormatInt(txnData.SourceAccountID, 10) + " to  " + strconv.FormatInt(txnData.DestinationAccountID, 10))
	}

	return u.transferRepo.SendMoney(newCtx, txnData)
}
