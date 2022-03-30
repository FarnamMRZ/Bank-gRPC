package app

import (
	"github.com/FarnamMRZ/Bank-gRPC/internal/service"
	pb "github.com/FarnamMRZ/Bank-gRPC/pkg"
)

type server struct {
	pb.UnimplementedBankServiceServer
	bankService     service.BankService
	customerService service.CustomerService
	accountService  service.AccountService
	transferService service.TransferService
}

func NewServer(bankService service.BankService, customerService service.CustomerService, accountService service.AccountService, transferService service.TransferService) *server {
	return &server{
		bankService:     bankService,
		customerService: customerService,
		accountService:  accountService,
		transferService: transferService,
	}
}
