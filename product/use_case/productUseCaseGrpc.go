package use_case

import (
	pb "github.com/kwul0208/common/api"
	models "github.com/kwul0208/product/model"
	"github.com/kwul0208/product/repository"
)

type ProductUseCaseGrpc interface {
	Create(grpcReq *pb.Product) (models.Product, error)
}

type productUseCaseGrpc struct {
	productRepository repository.ProductRepository
}

func NewProductUseCaseGrpc(productRepository repository.ProductRepository) *productUseCaseGrpc {
	return &productUseCaseGrpc{productRepository}
}

func (pu *productUseCaseGrpc) Create(grpcReq *pb.Product) (models.Product, error) {

	// Convert the gRPC product request to the models.Product type
	product := models.Product{
		ProductName: grpcReq.Name,
		Description: grpcReq.Description,
	}

	// Use the repository to save the product
	newProduct, err := pu.productRepository.Create(product)
	if err != nil {
		return models.Product{}, err // Return error if failed
	}

	// Return the newly created product and no error
	return newProduct, nil
}
