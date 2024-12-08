package main

import (
	"log"
	"net"

	"github.com/kwul0208/common"
	"github.com/kwul0208/user/handler"
	initializers "github.com/kwul0208/user/initializer"
	"github.com/kwul0208/user/repository"
	"github.com/kwul0208/user/use_case"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2001")
)

func main() {
	grpcServer := grpc.NewServer()

	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listern: %v.", err)
	}
	defer listen.Close()

	initializers.LoadEnvVariables()
	db := initializers.ConnectDatabase()

	repo := repository.NewAuthRepository(db)
	uc := use_case.NewAuthUseCaseGrpc(repo)
	handler.NewAuthGrpcHandler(grpcServer, uc)

	log.Printf("GRPC user server started at %v", grpcAddr)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
