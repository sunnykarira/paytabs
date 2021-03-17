package repository

import (
	"context"

	"github.com/paytabs/bank/model"
)

type (
	TransferRepositoryParams struct {
		AccountsData map[int64]*AccountDetails
	}

	transferRepo struct {
		AccountsData map[int64]*AccountDetails
	}
)

func NewTransferRepository(params TransferRepositoryParams) TransferRepository {

	return &transferRepo{
		AccountsData: params.AccountsData,
	}
}

func (t *transferRepo) SendMoney(ctx context.Context, txnData model.TransactionData) (success bool, err error) {

	s, ok := t.AccountsData[txnData.SourceAccountID]
	if !ok || s == nil {
		return false, ERR_INVALID_ACCOUNT_ID
	}

	d, ok := t.AccountsData[txnData.DestinationAccountID]
	if !ok || d == nil {
		return false, ERR_INVALID_ACCOUNT_ID
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	d.mu.Lock()
	defer d.mu.Unlock()
	s.data.Balance -= txnData.Amount
	d.data.Balance += txnData.Amount
	return true, nil
}
