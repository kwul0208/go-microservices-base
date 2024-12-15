package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kwul0208/user/pkg/config"
	"github.com/kwul0208/user/pkg/db"
	"github.com/kwul0208/user/pkg/handler"
	"github.com/kwul0208/user/pkg/pb"
	"github.com/kwul0208/user/pkg/repository"
	"github.com/kwul0208/user/pkg/use_case"
	"github.com/kwul0208/user/pkg/utils"
	"google.golang.org/grpc"
)

func main() {

	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed load config", err)
	}

	h := db.Init(c)

	jwt := utils.JWTWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed at listening : ", err)
	}

	fmt.Println("User scv on ", c.Port)

	repo := repository.NewAuthRepository(h)
	uc := use_case.NewAuthUseCaseGrpc(repo)
	// handler.NewAuthGrpcHandler(grpcServer, uc)

	s := handler.Server{
		H:        h,
		Jwt:      jwt,
		User_guc: uc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve : ", err)
	}

}
