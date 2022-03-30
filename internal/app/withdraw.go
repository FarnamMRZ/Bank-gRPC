package app

import (
	"context"
	pb "github.com/FarnamMRZ/Bank-gRPC/pkg"
)

func (s *server) Withdraw(ctx context.Context, in *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	userName := in.GetCustomerUserName()
	accountNumber := in.GetAccountNumber()
	amount := in.GetAmount()

	err := s.transferService.Withdraw(userName, accountNumber, amount)
	if err != nil {
		return nil, err
	}
	return &pb.WithdrawResponse{}, nil
}
