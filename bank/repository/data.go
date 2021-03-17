package repository

import (
	"context"
	"errors"
	"github.com/paytabs/bank/model"
	"strconv"
)

type (
	DataRepositoryParams struct {
		AccountsData map[int64]*AccountDetails
	}

	dataRepo struct {
		accountsData map[int64]*AccountDetails
	}
)

func NewDataRepository(params DataRepositoryParams) DataRepository {

	return &dataRepo{
		accountsData: params.AccountsData,
	}
}

func (a *dataRepo) FetchData(ctx context.Context, accountID int64) (accountDetails model.Account, err error) {

	if accountID <= 0 {
		return model.Account{}, errors.New("invalid account id " + strconv.FormatInt(accountID, 10))
	}

	v, ok := a.accountsData[accountID]
	if !ok || v == nil {
		return model.Account{}, errors.New("account does not exist " + strconv.FormatInt(accountID, 10))
	}
	v.mu.RLock()
	defer v.mu.RUnlock()
	select {
	case <-ctx.Done():
		return model.Account{}, ERR_REQUEST_FAILED_DUE_TO_TIMEOUT
	default:
	}
	return v.data, nil

}
