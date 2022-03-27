package app

import (
	pb "bank/pkg"
	"context"
)

func (s *server) CreateCustomer(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	userName := in.GetUserName()
	initMoney := in.GetInitMoney()

	err := s.customerService.CreateCustomer(userName, initMoney)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCustomerResponse{}, nil
}
