package app

import (
	"context"
	pb "github.com/FarnamMRZ/Bank-gRPC/pkg"
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
