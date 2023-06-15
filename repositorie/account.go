package repositorie

type Account struct {
	AccountID    int     `db:"account_id"`
	CustomerID   int     `db:"customer_id"`
	Opening_date string  `db:"opening_date"`
	Account_type string  `db:"account_type"`
	Amount       float64 `db:"amount"`
	Status       int     `db:"status"`
}

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
}
