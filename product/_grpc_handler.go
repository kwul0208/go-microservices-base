package main

import (
	"context"
	"log"

	pb "github.com/kwul0208/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedProductServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterProductServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateProduct(ctx context.Context, pr *pb.CreateProductRequest) (*pb.Product, error) {
	log.Printf("new product recive! (product) %v", pr)
	log.Print(pr)
	p := &pb.Product{
		ID: "1",
	}
	return p, nil
}
