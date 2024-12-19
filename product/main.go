package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kwul0208/product/pkg/config"
	"github.com/kwul0208/product/pkg/db"
	"github.com/kwul0208/product/pkg/handler"
	"github.com/kwul0208/product/pkg/pb"
	"github.com/kwul0208/product/pkg/repository"
	"github.com/kwul0208/product/pkg/use_case"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed load config", err)
	}

	h := db.Init(c)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed at listening: ", err)
	}

	fmt.Println("Product SVC on ", c.Port)

	repo := repository.NewProductRepository(h)
	uc := use_case.NewProductUseCaseGrpc(repo)

	s := handler.Server{
		H:           h,
		Product_guc: uc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve : ", err)
	}

}
