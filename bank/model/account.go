package model

type AccountStatus int

const (
	AccountStatusBlocked AccountStatus = 0
	AccountStatusActive  AccountStatus = 1
)

type Account struct {
	ID            int64
	Balance       float64
	Location      string
	AccountStatus AccountStatus
}

