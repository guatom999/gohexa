package service

import (
	"gohexa/errs"
	"gohexa/logs"
	"gohexa/repositorie"
	"time"
)

type accountService struct {
	accRepo repositorie.AccountRepository
}

func NewAccountService(accRepo repositorie.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*NewAccountResponse, error) {
	if request.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5000")
	}

	account := repositorie.Account{
		CustomerID:   customerID,
		Opening_date: time.Now().Format("2006-1-2 15:04:05"),
		Account_type: request.Account_type,
		Amount:       request.Amount,
		Status:       1,
	}

	newAcc, err := s.accRepo.Create(account)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnExpectedError()
	}

	response := NewAccountResponse{
		AccountID:    newAcc.AccountID,
		Opening_date: newAcc.Opening_date,
		Account_type: newAcc.Account_type,
		Amount:       newAcc.Amount,
		Status:       newAcc.Status,
	}
	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]NewAccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)

	if err != nil {
		logs.Error(err)
		return nil, err
	}

	responses := []NewAccountResponse{}
	for _, account := range accounts {
		responses = append(responses, NewAccountResponse{
			AccountID:    account.AccountID,
			Opening_date: account.Opening_date,
			Account_type: account.Account_type,
			Amount:       account.Amount,
			Status:       account.Status,
		})
	}
	return responses, nil
}
