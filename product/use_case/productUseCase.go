package use_case

import (
	models "github.com/kwul0208/product/model"
	"github.com/kwul0208/product/repository"
	"github.com/kwul0208/product/request"
)

type ProductUseCase interface {
	GetAll() ([]models.Product, error)
	GetById(Id int) (models.Product, error)
	Create(request request.ProductRequest) (models.Product, error)
	Update(Id int, request request.ProductRequest) (models.Product, error)
	Delete(Id int) (models.Product, error)
}

type productUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(productRepository repository.ProductRepository) *productUseCase {
	return &productUseCase{productRepository}
}

func (pu *productUseCase) GetAll() ([]models.Product, error) {
	products, err := pu.productRepository.GetAll()
	return products, err
}

func (pu *productUseCase) GetById(Id int) (models.Product, error) {
	product, err := pu.productRepository.GetById(Id)
	return product, err
}

func (pu *productUseCase) Create(productRequest request.ProductRequest) (models.Product, error) {
	product := models.Product{
		Id:          productRequest.Id,
		ProductName: productRequest.ProductName,
		Description: productRequest.Description,
	}

	newProduct, err := pu.productRepository.Create(product)

	return newProduct, err
}

func (pu *productUseCase) Update(Id int, productRequest request.ProductRequest) (models.Product, error) {
	product := models.Product{
		ProductName: productRequest.ProductName,
		Description: productRequest.Description,
	}

	updatedProduct, err := pu.productRepository.Update(Id, product)

	return updatedProduct, err
}

func (pu *productUseCase) Delete(Id int) (models.Product, error) {
	product, err := pu.productRepository.GetById(Id)
	deletedProduct, err := pu.productRepository.Delete(product)

	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusOK, gin.H{
	// 		"status": "error",
	// 		"message": err.Error(),
	// 	})
	// }
	return deletedProduct, err
}
