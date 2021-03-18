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

	if txnData.SourceAccountID == txnData.DestinationAccountID {
		return false, errors.New("cannot send money to same account")
	}

	var (
		srcData  model.Account
		destData model.Account
	)
	var (
		srcErr  error
		destErr error
	)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		srcData, srcErr = u.dataRepo.FetchData(ctx, txnData.SourceAccountID)
	}()

	go func() {
		defer wg.Done()
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
		return false, errors.New("accountID " + strconv.FormatInt(txnData.SourceAccountID, 10) + " blocked")
	}

	if destData.AccountStatus == model.AccountStatusBlocked {
		return false, errors.New("accountID " + strconv.FormatInt(txnData.DestinationAccountID, 10) + " blocked ")
	}

	if srcData.Balance-txnData.Amount < 0 {
		return false, errors.New("not enough balance in account to send money from " + strconv.FormatInt(txnData.SourceAccountID, 10) + " to  " + strconv.FormatInt(txnData.DestinationAccountID, 10))
	}

	return u.transferRepo.SendMoney(ctx, txnData)
}
