package usecase

import (
	"context"
	"errors"
	"github.com/paytabs/bank/model"
	"strconv"
	"sync"
)

func (u *usecase) FetchData(ctx context.Context, accountID int64) (accountDetails model.Account, err error) {

	return u.dataRepo.FetchData(ctx, accountID)
}

func (u *usecase) CreateAccount(ctx context.Context, accountData model.Account) (success bool, err error) {

	return u.accountsRepo.CreateAccount(ctx, accountData)
}

func (u *usecase) AddMoney(ctx context.Context, accountID int64, money float64) (success bool, err error) {

	return u.accountsRepo.AddMoney(ctx, accountID, money)
}

func (u *usecase) SendMoney(ctx context.Context, txnData model.TransactionData) (success bool, err error) {

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
		srcData, srcErr = u.dataRepo.FetchData(ctx, txnData.SourceAccountID)
	}()

	go func() {
		destData, destErr = u.dataRepo.FetchData(ctx, txnData.DestinationAccountID)
	}()
	wg.Wait()

	if srcErr != nil {
		return false, srcErr
	}
	if destErr != nil {
		return false, destErr
	}

	if srcData.AccountStatus == model.AccountStatusBlocked {
		return false, errors.New("account blocked " + strconv.FormatInt(txnData.SourceAccountID, 10))
	}

	if destData.AccountStatus == model.AccountStatusBlocked {
		return false, errors.New("account blocked " + strconv.FormatInt(txnData.DestinationAccountID, 10))
	}

	if srcData.Balance-txnData.Amount < 0 {
		return false, errors.New("not enough balance in account to send money from " + strconv.FormatInt(txnData.SourceAccountID, 10) + " to  " + strconv.FormatInt(txnData.DestinationAccountID, 10))
	}

	return u.transferRepo.SendMoney(ctx, txnData)
}
