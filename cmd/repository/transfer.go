package repository

import (
	"context"
	"github.com/paytabs/cmd/model"
)

type (
	TransferRepositoryParams struct{}

	transferRepo struct{}
)

func NewTransferRepository(params TransferRepositoryParams) TransferRepository {

	return &transferRepo{}
}

func (t *transferRepo) SendMoney(ctx context.Context, txnData model.TransactionData) (success bool, err error) {
	return false, nil
}
