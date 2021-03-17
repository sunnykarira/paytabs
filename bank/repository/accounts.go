package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/paytabs/bank/model"
)

type (
	AccountDetails struct {
		data model.Account
		mu   *sync.RWMutex
	}

	AccountsRepositoryParams struct {
		AccountsData map[int64]*AccountDetails
	}

	accountsRepo struct {
		accountsData map[int64]*AccountDetails
	}
)

func NewAccountRepository(params AccountsRepositoryParams) AccountsRepository {

	return &accountsRepo{
		accountsData: params.AccountsData,
	}
}

var (
	ERR_INVALID_ACCOUNT_ID            = errors.New("invalid account id")
	ERR_ACCOUNT_DOES_NOT_EXIST        = errors.New("account does not exist")
	ERR_ACCOUNT_ALREADY_EXIST         = errors.New("account already exist")
	ERR_INVALID_MONEY_TO_ADD          = errors.New("invalid money to add")
	ERR_REQUEST_FAILED_DUE_TO_TIMEOUT = errors.New("failed due to timeout")
	ERR_ACCOUNT_BLOCKED               = errors.New("account blocked")
)

func (a *accountsRepo) CreateAccount(ctx context.Context, accountData model.Account) (success bool, err error) {

	if accountData.ID <= 0 {
		return false, ERR_INVALID_ACCOUNT_ID
	}
	_, ok := a.accountsData[accountData.ID]
	if ok {
		return false, ERR_ACCOUNT_ALREADY_EXIST
	}

	a.accountsData[accountData.ID] = &AccountDetails{
		data: accountData,
		mu:   &sync.RWMutex{},
	}
	select {
	case <-ctx.Done():
		return false, ERR_REQUEST_FAILED_DUE_TO_TIMEOUT
	default:
	}
	return true, nil
}
func (a *accountsRepo) AddMoney(ctx context.Context, accountID int64, money float64) (success bool, err error) {

	if accountID <= 0 {
		return false, ERR_INVALID_ACCOUNT_ID
	}

	if money <= 0 {
		return false, ERR_INVALID_MONEY_TO_ADD
	}

	v, ok := a.accountsData[accountID]
	if !ok || v == nil {
		return false, ERR_ACCOUNT_DOES_NOT_EXIST
	}

	if v.data.AccountStatus == model.AccountStatusBlocked {
		return false, ERR_ACCOUNT_BLOCKED
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	select {
	case <-ctx.Done():
		return false, ERR_REQUEST_FAILED_DUE_TO_TIMEOUT
	default:
	}

	v.data.Balance += money
	return true, nil
}
