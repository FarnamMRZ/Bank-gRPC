package app

import (
	"context"
	pb "github.com/FarnamMRZ/Bank-gRPC/pkg"
)

func (s *server) CreateBank(ctx context.Context, in *pb.CreateBankRequest) (*pb.CreateBankResponse, error) {
	bankName := in.GetName()
	err := s.bankService.CreateBank(bankName)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBankResponse{}, nil
}
