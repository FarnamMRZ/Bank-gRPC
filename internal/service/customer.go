package service

import (
	"github.com/FarnamMRZ/Bank-gRPC/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerService interface {
	CreateCustomer(userName string, initMoney int64) error
}

type customerService struct {
	dao repository.DAO
}

func NewCustomerService(dao repository.DAO) CustomerService {
	return &customerService{dao}
}

func (cs *customerService) CreateCustomer(userName string, initMoney int64) error {
	doesExist, err := cs.dao.NewCustomerQuery().CustomerExist(userName)
	if err != nil {
		return err
	}
	if doesExist {
		return status.Errorf(codes.AlreadyExists, "az in moshtari chand ta mikhay lashi?")
	}
	return cs.dao.NewCustomerQuery().CreateCustomer(userName, initMoney)
}
