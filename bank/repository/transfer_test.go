package repository

import (
	"context"
	"github.com/paytabs/bank/model"
	"sync"
	"testing"
)

func Test_SendMoneyt(t *testing.T) {

	type args struct {
		txnData model.TransactionData
		ctx         context.Context
	}

	tests := []struct {
		name        string
		args        args
		wantSuccess bool
		wantErr     bool
		f           func(mp map[int64]*AccountDetails)
		m           map[int64]*AccountDetails
	}{
		{
			name: "success",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID: 10,
					DestinationAccountID: 11,
					Amount: 10,
				},
				ctx: context.Background(),
			},
			f:           func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{
					mu : &sync.RWMutex{},
					data: model.Account{
						ID: 10,
						Location: "test",
						Balance: 20,
						AccountStatus: model.AccountStatusActive,
					},
				}
				mp[11] = &AccountDetails{
					mu : &sync.RWMutex{},
					data: model.Account{
						ID: 11,
						Location: "test",
						Balance: 20,
						AccountStatus: model.AccountStatusActive,
					},
				}
			},
			wantSuccess: true,
		},
		{
			name: "invalid source account",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID: 10,
					DestinationAccountID: 11,
					Amount: 10,
				},
				ctx: context.Background(),
			},
			f:           func(mp map[int64]*AccountDetails) {
				mp[11] = &AccountDetails{
					mu : &sync.RWMutex{},
					data: model.Account{
						ID: 11,
						Location: "test",
						Balance: 20,
						AccountStatus: model.AccountStatusActive,
					},
				}
			},
			wantErr: true,
		},
		{
			name: "invalid destination account",
			args: args{
				txnData: model.TransactionData{
					SourceAccountID: 10,
					DestinationAccountID: 11,
					Amount: 10,
				},
				ctx: context.Background(),
			},
			f:           func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{
					mu : &sync.RWMutex{},
					data: model.Account{
						ID: 10,
						Location: "test",
						Balance: 20,
						AccountStatus: model.AccountStatusActive,
					},
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := make(map[int64]*AccountDetails)
			tt.f(mp)
			s, r := (&transferRepo{
				accountsData: mp,
			}).SendMoney(tt.args.ctx, tt.args.txnData)
			if (r != nil) != tt.wantErr {
				t.Fatalf("unexpected error %+v", r)
			}
			if s != tt.wantSuccess {
				t.Fatalf("unexpected success got %+v want %+v", s, tt.wantSuccess)
			}
		})
	}
}

