package repository

import (
	"context"
	"github.com/paytabs/bank/model"
	"sync"
	"testing"
)

func Test_CreateAccount(t *testing.T) {

	type args struct {
		accountData model.Account
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
				accountData: model.Account{
					ID:            10,
					Balance:       20.22,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				},
				ctx: context.Background(),
			},
			f:           func(mp map[int64]*AccountDetails) {},
			wantSuccess: true,
		},
		{
			name: "invalid account id",
			args: args{
				accountData: model.Account{
					ID:            0,
					Balance:       20.22,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				},
				ctx: context.Background(),
			},
			f:       func(mp map[int64]*AccountDetails) {},
			wantErr: true,
		},
		{
			name: "account already present",
			args: args{
				accountData: model.Account{
					ID:            10,
					Balance:       20.22,
					Location:      "test",
					AccountStatus: model.AccountStatusActive,
				},
				ctx: context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := make(map[int64]*AccountDetails)
			tt.f(mp)
			s, r := (&accountsRepo{
				accountsData: mp,
			}).CreateAccount(tt.args.ctx, tt.args.accountData)
			if (r != nil) != tt.wantErr {
				t.Fatalf("unexpected error %+v", r)
			}
			if s != tt.wantSuccess {
				t.Fatalf("unexpected success got %+v want %+v", s, tt.wantSuccess)
			}
		})
	}
}

func Test_AddMoney(t *testing.T) {

	type args struct {
		accountID int64
		money     float64
		ctx       context.Context
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
				accountID: 10,
				money:     20,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{
					data: model.Account{
						ID:            10,
						Balance:       10,
						Location:      "test",
						AccountStatus: model.AccountStatusActive,
					},
					mu: &sync.RWMutex{},
				}
			},
			wantSuccess: true,
		},
		{
			name: "invalid account id",
			args: args{
				accountID: 0,
				money:     20,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{
					data: model.Account{
						ID:            10,
						Balance:       10,
						Location:      "test",
						AccountStatus: model.AccountStatusActive,
					},
					mu: &sync.RWMutex{},
				}
			},
			wantErr: true,
		},
		{
			name: "invalid money",
			args: args{
				accountID: 10,
				money:     -10,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{
					data: model.Account{
						ID:            10,
						Balance:       10,
						Location:      "test",
						AccountStatus: model.AccountStatusActive,
					},
					mu: &sync.RWMutex{},
				}
			},
			wantErr: true,
		},
		{
			name: "account does not exists",
			args: args{
				accountID: 10,
				money:     10,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {

			},
			wantErr: true,
		},
		{
			name: "account blocked",
			args: args{
				accountID: 10,
				money:     10,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{
					data: model.Account{
						ID:            10,
						Balance:       10,
						Location:      "test",
						AccountStatus: model.AccountStatusBlocked,
					},
					mu: &sync.RWMutex{},
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := make(map[int64]*AccountDetails)
			tt.f(mp)
			s, r := (&accountsRepo{
				accountsData: mp,
			}).AddMoney(tt.args.ctx, tt.args.accountID, tt.args.money)
			if (r != nil) != tt.wantErr {
				t.Fatalf("unexpected error %+v", r)
			}
			if s != tt.wantSuccess {
				t.Fatalf("unexpected success got %+v want %+v", s, tt.wantSuccess)
			}
		})
	}
}
