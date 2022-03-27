package service

import (
	"bank/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransferService interface {
	Withdraw(userName, accountNumber string, amount int64) error
}

type transferService struct {
	dao repository.DAO
}

func NewTransferService(dao repository.DAO) TransferService {
	return &transferService{dao}
}

func (ts *transferService) Withdraw(userName, accountNumber string, amount int64) error {
	doesExist, err := ts.dao.NewCustomerQuery().CustomerExist(userName)
	if err != nil {
		return err
	}
	if !doesExist {
		return status.Errorf(codes.NotFound, "codom mashtari dash?")
	}

	doesExist, err = ts.dao.NewAccountQuery().AccountExist(userName, accountNumber)
	if err != nil {
		return err
	}
	if !doesExist {
		return status.Errorf(codes.NotFound, "chizi zadi?")
	}

	accountAmount, err := ts.dao.NewAccountQuery().GetAccountAmount(userName, accountNumber)
	if err != nil {
		return err
	}

	if *accountAmount < amount {
		return status.Errorf(codes.FailedPrecondition, "boro baba pool nadari!")
	}

	err = ts.dao.NewAccountQuery().UpdateAccount(userName, accountNumber, *accountAmount-amount)
	if err != nil {
		return err
	}

	safeAmount, err := ts.dao.NewCustomerQuery().GetSafeAmount(userName)
	if err != nil {
		return err
	}

	err = ts.dao.NewCustomerQuery().UpdateSafeAmount(userName, *safeAmount+amount)
	if err != nil {
		return err
	}
	return nil
}
