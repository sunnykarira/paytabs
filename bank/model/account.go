package model

type AccountStatus int

const (
	AccountStatusBlocked AccountStatus = 0
	AccountStatusActive  AccountStatus = 1
)

type Account struct {
	ID            int64         `json:"id"`
	Balance       float64       `json:"balance"`
	Location      string        `json:"location"`
	AccountStatus AccountStatus `json:"account_status"`
}
