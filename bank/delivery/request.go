package delivery

type AccountDetailsRequest struct {
	ID int64 `json:"account_id"`
}

type CreateAccountRequest struct {
	ID            int64   `json:"account_id"`
	Balance       float64 `json:"balance"`
	Location      string  `json:"location"`
	AccountStatus int     `json:"account_status"`
}

type AddMoneyRequest struct {
	ID     int64   `json:"account_id"`
	Amount float64 `json:"amount"`
}

type SendMoneyRequest struct {
	SourceAccountID      int64   `json:"source_account_id"`
	DestinationAccountID int64   `json:"destination_account_id"`
	Amount               float64 `json:"amount"`
}
