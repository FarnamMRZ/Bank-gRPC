package app

import (
	pb "bank/pkg"
	"context"
)

func (s *server) CreateBank(ctx context.Context, in *pb.CreateBankRequest) (*pb.CreateBankResponse, error) {
	bankName := in.GetName()
	err := s.bankService.CreateBank(bankName)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBankResponse{}, nil
}
