package repository

import (
	"context"
	"github.com/paytabs/bank/model"
	"sync"
	"testing"
)

func Test_FetchData(t *testing.T) {

	type args struct {
		accountID int64
		ctx       context.Context
	}

	tests := []struct {
		name          string
		args          args
		wantAccountID int64
		wantErr       bool
		f             func(mp map[int64]*AccountDetails)
		m             map[int64]*AccountDetails
	}{
		{
			name: "success",
			args: args{
				accountID: 10,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {
				mp[10] = &AccountDetails{
					data: model.Account{
						ID: 10,
					},
					mu: &sync.RWMutex{},
				}
			},
			wantAccountID: 10,
		},
		{
			name: "invalid account id",
			args: args{
				accountID: 0,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {

			},
			wantErr:       true,
			wantAccountID: 0,
		},
		{
			name: "account does not exists" ,
			args: args{
				accountID: 10,
				ctx:       context.Background(),
			},
			f: func(mp map[int64]*AccountDetails) {

			},
			wantErr:       true,
			wantAccountID: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := make(map[int64]*AccountDetails)
			tt.f(mp)
			s, r := (&dataRepo{
				accountsData: mp,
			}).FetchData(tt.args.ctx, tt.args.accountID)
			if (r != nil) != tt.wantErr {
				t.Fatalf("unexpected error %+v", r)
			}
			if s.ID != tt.wantAccountID {
				t.Fatalf("unexpected success got %+v want %+v", s, tt.wantAccountID)
			}
		})
	}
}
