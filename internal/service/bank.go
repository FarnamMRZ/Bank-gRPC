package service

import (
	"bank/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BankService interface {
	CreateBank(bankName string) error
}

type bankService struct {
	dao repository.DAO
}

func NewBankService(dao repository.DAO) BankService {
	return &bankService{dao}
}

func (bs *bankService) CreateBank(bankName string) error {
	doesExist, err := bs.dao.NewBankQuery().BankExist(bankName)
	if err != nil {
		return err
	}
	if doesExist {
		return status.Errorf(codes.AlreadyExists, "az in bank chand ta mikhay lashi?")
	}
	return bs.dao.NewBankQuery().CreateBank(bankName)
}
