package handler

import (
	"context"
	"log"

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

func (h *grpcHandler) GetProducts(ctx context.Context, pr *pb.Empty) (*pb.Products, error) {
	log.Print("p")
	products, err := h.product_guc.Get()

	if err != nil {
		log.Printf("Error getting product: %v", err)
		return nil, err
	}

	var pbProducts []*pb.Product
	for _, p := range products {

		pbProduct := &pb.Product{
			ID:          p.Id,
			Name:        p.ProductName,
			Description: p.Description,
		}
		pbProducts = append(pbProducts, pbProduct)
	}

	return &pb.Products{
		Products: pbProducts, // Return the slice of products
	}, nil
}

func (h *grpcHandler) GetProductById(ctx context.Context, pr *pb.ProductID) (*pb.Product, error) {
	product, err := h.product_guc.GetById(pr.ID)

	productReturn := &pb.Product{
		ID:          product.Id,
		Name:        product.ProductName,
		Description: product.Description,
	}

	if err != nil {
		log.Print("error get by id")
	}

	return productReturn, err
}

func (h *grpcHandler) CreateProduct(ctx context.Context, pr *pb.CreateProductRequest) (*pb.Product, error) {

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

	// id, _ := strconv.Atoi(pr.ID)

	_, err := h.product_guc.Update(int(pr.ID), p)

	if err != nil {
		log.Printf("Error update")
	}

	return p, nil
}
