package service

import (
	"bank/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type AccountService interface {
	CreateAccount(userName, bankName string, initDeposit int64) error
}

type accountService struct {
	dao repository.DAO
}

func NewAccountService(dao repository.DAO) AccountService {
	return &accountService{dao}
}

func (as *accountService) CreateAccount(userName, bankName string, initDeposit int64) error {
	doesExist, err := as.dao.NewCustomerQuery().CustomerExist(userName)
	if err != nil {
		return err
	}
	if !doesExist {
		return status.Errorf(codes.NotFound, "codom mashtari dash?")
	}

	doesExist, err = as.dao.NewBankQuery().BankExist(bankName)
	if !doesExist {
		return status.Errorf(codes.NotFound, "in bank koddom keshvare?")
	}

	safeAmount, err := as.dao.NewCustomerQuery().GetSafeAmount(userName)
	if err != nil {
		return err
	}

	if *safeAmount < initDeposit {
		return status.Errorf(codes.FailedPrecondition, "boro baba pool nadari!")
	}

	number, err := as.dao.NewCustomerQuery().GetNumberOfAccounts(userName)
	if err != nil {
		return err
	}

	newNumber := strconv.Itoa(int(*number + 1))
	err = as.dao.NewAccountQuery().CreateAccount(userName, bankName, newNumber, initDeposit)
	if err != nil {
		return err
	}

	err = as.dao.NewCustomerQuery().AddAccount(userName)
	if err != nil {
		return err
	}

	err = as.dao.NewCustomerQuery().UpdateSafeAmount(userName, *safeAmount-initDeposit)
	if err != nil {
		return err
	}

	return nil
}
