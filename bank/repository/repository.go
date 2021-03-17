package repository

import (
	"context"

	"github.com/paytabs/bank/model"
)

type (
	DataRepository interface {
		FetchData(ctx context.Context, accountID int64) (accountDetails model.Account, err error)
	}

	AccountsRepository interface {
		CreateAccount(ctx context.Context, accountData model.Account) (success bool, err error)
		AddMoney(ctx context.Context, accountID int64, money float64) (success bool, err error)
	}

	TransferRepository interface {
		SendMoney(ctx context.Context, txnData model.TransactionData) (success bool, err error)
	}
)
