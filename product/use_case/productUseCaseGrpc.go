package use_case

import (
	pb "github.com/kwul0208/common/api"
	models "github.com/kwul0208/product/model"
	"github.com/kwul0208/product/repository"
)

type ProductUseCaseGrpc interface {
	Get() ([]models.Product, error)
	GetById(Id int64) (models.Product, error)
	Create(grpcReq *pb.Product) (models.Product, error)
	Update(Id int, grpcReq *pb.Product) (models.Product, error)
	Delete(Id int64) (models.Product, error)
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

func (pu *productUseCaseGrpc) Get() ([]models.Product, error) {
	products, err := pu.productRepository.GetAll()

	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
}

func (pu *productUseCaseGrpc) GetById(Id int64) (models.Product, error) {
	product, err := pu.productRepository.GetById(int(Id))

	if err != nil {
		return models.Product{}, err
	}

	return product, err
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

func (pu *productUseCaseGrpc) Update(Id int, grpcReq *pb.Product) (models.Product, error) {
	product := models.Product{
		ProductName: grpcReq.Name,
		Description: grpcReq.Description,
	}

	updatedProduct, err := pu.productRepository.Update(Id, product)

	return updatedProduct, err

}

func (pu *productUseCaseGrpc) Delete(Id int64) (models.Product, error) {
	product, err := pu.productRepository.GetById(int(Id))
	if err != nil {
		return models.Product{}, err
	}
	deletedProduct, _ := pu.productRepository.Delete(product)

	return deletedProduct, nil
}
