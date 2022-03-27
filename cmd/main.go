package main

import (
	"bank/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"

	"bank/internal/app"
	"bank/internal/repository"
	pb "bank/pkg"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Faild to listen err: %v", err)
	}

	// Register dao
	s, err := repository.NewDB()
	if err != nil {
		log.Fatalf("Faild to connect to database: %v", err)
	}

	dao := repository.NewDAO(s)

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
