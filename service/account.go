package service

type NewAccountRequest struct {
	Account_type string  `json:"account_type"`
	Amount       float64 `json:"amount"`
}

type NewAccountResponse struct {
	AccountID    int     `json:"account_id"`
	Opening_date string  `json:"opening_date"`
	Account_type string  `json:"account_type"`
	Amount       float64 `json:"amount"`
	Status       int     `json:"status"`
}

type AccountService interface {
	NewAccount(int, NewAccountRequest) (*NewAccountResponse, error)
	GetAccounts(int) ([]NewAccountResponse, error)
}
