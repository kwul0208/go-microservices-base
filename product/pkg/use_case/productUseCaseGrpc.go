package use_case

import (
	models "github.com/kwul0208/product/pkg/model"
	"github.com/kwul0208/product/pkg/pb"
	"github.com/kwul0208/product/pkg/repository"
)

type ProductUseCaseGrpc interface {
	GetAll(request *pb.FindAllRequest) ([]models.Product, error)
	GetOne(request *pb.FindOneRequest) (models.Product, error)
	CreateProduct(request *pb.CreateProductRequest) (models.Product, error)
}

type productUseCaseGrpc struct {
	productRepository repository.ProductRepository
}

// // Update implements ProductUseCaseGrpc.
// func (pu *productUseCaseGrpc) Update(Id int64, grpcReq *pb.Product) (models.Product, error) {
// 	panic("unimplemented")
// }

func NewProductUseCaseGrpc(productRepository repository.ProductRepository) *productUseCaseGrpc {
	return &productUseCaseGrpc{productRepository}
}

func (pu *productUseCaseGrpc) GetAll(grpc *pb.FindAllRequest) ([]models.Product, error) {
	products, err := pu.productRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pu *productUseCaseGrpc) GetOne(grpcReq *pb.FindOneRequest) (models.Product, error) {
	product, err := pu.productRepository.GetOne(int(grpcReq.Id))

	if err != nil {
		return models.Product{}, err
	}

	return product, err
}

func (pu *productUseCaseGrpc) CreateProduct(grpcReq *pb.CreateProductRequest) (models.Product, error) {

	// Convert the gRPC product request to the models.Product type
	product := models.Product{
		Name:  grpcReq.Name,
		Price: grpcReq.Price,
		Stock: grpcReq.Stock,
	}

	// Use the repository to save the product
	newProduct, err := pu.productRepository.CreateProduct(product)
	if err != nil {
		return models.Product{}, err // Return error if failed
	}

	// Return the newly created product and no error
	return newProduct, nil
}
