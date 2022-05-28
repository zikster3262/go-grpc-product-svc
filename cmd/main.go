package main

import (
	"fmt"
	"log"
	"net"

	"github.com/zikster3262/go-grpc-product-svc/pkg/config"
	"github.com/zikster3262/go-grpc-product-svc/pkg/db"
	pb "github.com/zikster3262/go-grpc-product-svc/pkg/pb"
	services "github.com/zikster3262/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Printf("Server Order service is running on port%v.", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
