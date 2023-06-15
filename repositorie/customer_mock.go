package repositorie

import "errors"

type CustomerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1, Name: "Bossza", Status: 2},
		{CustomerID: 2, Name: "TesT", Status: 3},
		{CustomerID: 3, Name: "Pita", Status: 4},
	}
	return CustomerRepositoryMock{customers: customers}
}

func (r CustomerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r CustomerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}

	}

	return nil, errors.New("Customer Not Found")

}
