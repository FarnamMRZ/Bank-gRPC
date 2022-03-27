package app

import (
	pb "bank/pkg"
	"context"
)

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	userName := in.GetCustomerUserName()
	bankName := in.GetBankName()
	initDeposit := in.GetInitDeposit()

	err := s.accountService.CreateAccount(userName, bankName, initDeposit)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{}, nil
}
