package handler

import (
	"context"
	"log"
	"strconv"

	pb "github.com/kwul0208/common/api"
	"github.com/kwul0208/product/use_case"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedProductServiceServer
	product_guc use_case.ProductUseCaseGrpc // Make sure this is initialized
}

// Modify constructor to accept the use case as a parameter
func NewGRPCHandler(grpcServer *grpc.Server, productUseCase use_case.ProductUseCaseGrpc) {
	handler := &grpcHandler{
		product_guc: productUseCase, // Initialize the use case
	}
	pb.RegisterProductServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateProduct(ctx context.Context, pr *pb.CreateProductRequest) (*pb.Product, error) {

	log.Print("create")
	// Convert the request to a product model
	p := &pb.Product{
		Name:        pr.ProductOnly.ProductName,
		Description: pr.ProductOnly.ProductDescription,
	}

	// Use the use case to create the product
	_, err := h.product_guc.Create(p)
	if err != nil {
		log.Printf("Error creating product: %v", err)
		return nil, err
	}

	return p, nil
}

func (h *grpcHandler) UpdateProduct(ctx context.Context, pr *pb.UpdateProductRequest) (*pb.Product, error) {
	log.Print("update")
	p := &pb.Product{
		Name:        pr.ProductOnly.ProductName,
		Description: pr.ProductOnly.ProductDescription,
	}

	id, _ := strconv.Atoi(pr.ID)

	_, err := h.product_guc.Update(id, p)

	if err != nil {
		log.Printf("Error update")
	}

	return p, nil
}
