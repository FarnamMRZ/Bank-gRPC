package main

import (
	"github.com/FarnamMRZ/Bank-gRPC/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/FarnamMRZ/Bank-gRPC/internal/app"
	"github.com/FarnamMRZ/Bank-gRPC/internal/repository"
	pb "github.com/FarnamMRZ/Bank-gRPC/pkg"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Faild to listen err: %v", err)
	}

	// Register dao
	c, err := repository.NewDB()
	if err != nil {
		log.Fatalf("Faild to connect to database: %v", err)
	}
	defer repository.CloseDB(c)

	dao := repository.NewDAO(c)

	// Register all services
	bankService := service.NewBankService(dao)
	customerService := service.NewCustomerService(dao)
	accountService := service.NewAccountService(dao)
	transferService := service.NewTransferService(dao)

	// register server
	grpcServer := grpc.NewServer()
	pb.RegisterBankServiceServer(grpcServer, app.NewServer(bankService, customerService, accountService, transferService))

	// serve server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Faild to serve the server: %v", err)
	}
}
