package repositorie

import "github.com/jmoiron/sqlx"

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id , name from techcoach.customers2"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil

}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id , name from techcoach.customers2 where customer_id =?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
