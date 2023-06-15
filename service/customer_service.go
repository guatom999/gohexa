package service

import (
	"database/sql"
	"errors"
	"gohexa/errs"
	"gohexa/logs"
	"gohexa/repositorie"
	"log"
)

type customerService struct {
	customerRepo repositorie.CustomerRepository
}

func NewCustomerService(customerRepo repositorie.CustomerRepository) customerService {
	return customerService{customerRepo: customerRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	custermers, err := s.customerRepo.GetAll()
	if err != nil {

		if err == sql.ErrNoRows {
			log.Println(err)
			logs.Error(err)
			return nil, errors.New("customer id not found")
		}

		log.Println(err)
		return nil, err
	}

	custReponses := []CustomerResponse{}
	for _, customer := range custermers {
		custReponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custReponses = append(custReponses, custReponse)
	}
	return custReponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {

	customer, err := s.customerRepo.GetById(id)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotfoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnExpectedError()

	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil

}
