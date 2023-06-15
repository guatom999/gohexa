package repositorie

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := "INSERT INTO account (customer_id , opening_date , account_type , amount , status) values ( ? , ? , ? , ? , ?)"
	res, err := r.db.Exec(
		query,
		acc.CustomerID,
		acc.Opening_date,
		acc.Account_type,
		acc.Amount,
		acc.Status,
	)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)

	return &acc, nil

}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query := "select account_id , customer_id , opening_date , account_type , amount , status from account where customer_id = ?"
	account := []Account{}
	err := r.db.Select(&account, query, customerID)
	if err != nil {
		return nil, err
	}
	return account, nil
}
