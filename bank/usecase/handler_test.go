package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/paytabs/bank/repository/mocks"
	"testing"

	"github.com/paytabs/bank/model"
)

func Test_SendMoney(t *testing.T) {

	type args struct {
		txnData model.TransactionData
		ctx     context.Context
	}

	tests := []struct {
		name        string
		args        args
		wantSuccess bool
		wantErr     bool
		f           func(accountRepo *mocks.MockAccountsRepository, transferRepo *mocks.MockTransferRepository, dataRepo *mocks.MockDataRepository)
	}{
		{
			name: "success",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID:      10,
					DestinationAccountID: 11,
					Amount:               10,
				},
				ctx: context.Background(),
			},
			f: func(accountRepo *mocks.MockAccountsRepository, transferRepo *mocks.MockTransferRepository, dataRepo *mocks.MockDataRepository) {

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(10)).Return(model.Account{
					ID:            10,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(11)).Return(model.Account{
					ID:            11,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

				transferRepo.EXPECT().SendMoney(gomock.Any(), model.TransactionData{
					SourceAccountID:      int64(10),
					DestinationAccountID: int64(11),
					Amount:               float64(10),
				}).Return(true, nil)
			},
			wantSuccess: true,
		},
		{
			name: "error source",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID:      10,
					DestinationAccountID: 11,
					Amount:               10,
				},
				ctx: context.Background(),
			},
			f: func(accountRepo *mocks.MockAccountsRepository, transferRepo *mocks.MockTransferRepository, dataRepo *mocks.MockDataRepository) {

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(10)).Return(model.Account{

				}, errors.New("error"))

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(11)).Return(model.Account{
					ID:            11,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

			},
			wantErr: true,
		},
		{
			name: "error dest",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID:      10,
					DestinationAccountID: 11,
					Amount:               10,
				},
				ctx: context.Background(),
			},
			f: func(accountRepo *mocks.MockAccountsRepository, transferRepo *mocks.MockTransferRepository, dataRepo *mocks.MockDataRepository) {

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(10)).Return(model.Account{
					ID:            10,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(11)).Return(model.Account{

				}, errors.New("error"))

			},
			wantErr: true,
		},
		{
			name: "source blocked",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID:      10,
					DestinationAccountID: 11,
					Amount:               10,
				},
				ctx: context.Background(),
			},
			f: func(accountRepo *mocks.MockAccountsRepository, transferRepo *mocks.MockTransferRepository, dataRepo *mocks.MockDataRepository) {

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(10)).Return(model.Account{
					ID:            10,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusBlocked,
				}, nil)

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(11)).Return(model.Account{
					ID:            11,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

			},
			wantErr: true,
		},
		{
			name: "dest blocked",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID:      10,
					DestinationAccountID: 11,
					Amount:               10,
				},
				ctx: context.Background(),
			},
			f: func(accountRepo *mocks.MockAccountsRepository, transferRepo *mocks.MockTransferRepository, dataRepo *mocks.MockDataRepository) {

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(10)).Return(model.Account{
					ID:            10,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(11)).Return(model.Account{
					ID:            11,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusBlocked,
				}, nil)

			},
			wantErr: true,
		},
		{
			name: "source not enough balance",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID:      10,
					DestinationAccountID: 11,
					Amount:               10,
				},
				ctx: context.Background(),
			},
			f: func(accountRepo *mocks.MockAccountsRepository, transferRepo *mocks.MockTransferRepository, dataRepo *mocks.MockDataRepository) {

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(10)).Return(model.Account{
					ID:            10,
					Balance:       2,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

				dataRepo.EXPECT().FetchData(gomock.Any(), int64(11)).Return(model.Account{
					ID:            11,
					Balance:       10,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				}, nil)

			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			accountRepo := mocks.NewMockAccountsRepository(ctrl)
			transferRepo := mocks.NewMockTransferRepository(ctrl)
			dataRepo := mocks.NewMockDataRepository(ctrl)
			tt.f(accountRepo, transferRepo, dataRepo)
			u := NewUsecase(UseCaseInit{
				TransferRepository: transferRepo,
				DataRepository:     dataRepo,
				AccountsRepository: accountRepo,
			})

			s, r := u.SendMoney(tt.args.ctx, tt.args.txnData)
			if (r != nil) != tt.wantErr {
				t.Fatalf("unexpected error %+v", r)
			}
			if s != tt.wantSuccess {
				t.Fatalf("unexpected success got %+v want %+v", s, tt.wantSuccess)
			}
		})
	}
}
